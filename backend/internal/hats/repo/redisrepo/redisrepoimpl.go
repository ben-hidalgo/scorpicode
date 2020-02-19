package redisrepo

import (
	"backend/internal/hats/repo"
	"errors"
	"fmt"

	"github.com/rs/xid"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// Prefix is part of the hat:<id> key
const Prefix = "hat"

// SetName is the key for the set of hats
const SetName = "hats"

// idkey returns the id and the key, creating a new one if inputID is empty
// e.g. idkey("123") returns id -> "123" and key -> "hat:123"
func idkey(inputID string) (id string, key string) {
	if inputID == "" {
		id = xid.New().String()
	} else {
		id = inputID
	}
	key = fmt.Sprintf("%s:%s", Prefix, id)
	return
}

// FindAll queries all records
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {

	// TODO: SMEMBERS has no pagination
	// values will be an array of strings
	values, err := redis.Values(r.Conn.Do(SMEMBERS, SetName))
	if err != nil {
		return
	}

	// TODO: v is a string, better naming and/or conversion
	for _, v := range values {
		// logrus.Printf("v=%s", v)

		_, key := idkey(fmt.Sprintf("%s", v))
		all, err := redis.Values(r.Conn.Do(HGETALL, key))
		if err != nil {
			return nil, err
		}

		var hat repo.HatMod
		err = redis.ScanStruct(all, &hat)
		if err != nil {
			return nil, err
		}

		hats = append(hats, &hat)
	}

	return hats, nil
}

// Save performs an upsert
func (r *Repo) Save(hm repo.HatMod) (*repo.HatMod, error) {

	if hm.ID == "" && hm.Version != 0 {
		return nil, repo.ErrVersionMismatch
	}

	id, key := idkey(hm.ID)

	mod := &repo.HatMod{
		ID:      id,
		Color:   hm.Color,
		Name:    hm.Name,
		Inches:  hm.Inches,
		Version: hm.Version + 1,
	}

	//TODO: if id is not populated, insert; populated created_at, updated_at and add a version for optimistic locking
	// add watch()

	if err := r.multi(); err != nil {
		logrus.Errorf("multi() err=%v", err)
		return nil, err
	}
	defer r.discard()

	if _, err := r.Conn.Do(HMSET, redis.Args{}.Add(key).AddFlat(mod)...); err != nil {
		logrus.Errorf("hmset err=%v", err)
		return nil, err
	}

	// set add hats <id>
	if _, err := r.Conn.Do(SADD, SetName, mod.ID); err != nil {
		logrus.Errorf("sadd err=%v", err)
		return nil, err
	}

	if err := r.exec(); err != nil {
		logrus.Errorf("exec() err=%v", err)
		return nil, err
	}

	return mod, nil
}

// Delete deletes the record if version matches; throws NotFound, VersionMismatch
func (r *Repo) Delete(id string, version int) error {

	// check for existence and version match
	mod, err := r.Find(id)
	if err != nil {
		return err
	}
	if mod == nil {
		return repo.ErrNotFound
	}
	if mod.Version != version {
		return repo.ErrVersionMismatch
	}

	// begin multi
	if err := r.multi(); err != nil {
		logrus.Errorf("Delete() multi() err=%v", err)
		return err
	}
	defer r.discard()

	// remove the id from the set (hats)
	_, err = r.Conn.Do(SREM, SetName, id)
	if err != nil {
		return err
	}

	_, key := idkey(id)

	// delete the key hat:<id>
	_, err = r.Conn.Do(DEL, key)
	if err != nil {
		return err
	}

	if err := r.exec(); err != nil {
		logrus.Errorf("exec() err=%v", err)
		return err
	}

	return nil
}

// Exists returns true if the record exists
func (r *Repo) Exists(id string) (bool, error) {

	// set is member
	// value is an integer: 1 or 0
	v, err := redis.Int(r.Conn.Do(SISMEMBER, SetName, id))
	if err != nil {
		return false, err
	}

	return v == 1, nil
}

// Find one; returns NotFound
func (r *Repo) Find(id string) (*repo.HatMod, error) {

	var mod repo.HatMod

	_, key := idkey(id)

	v, err := redis.Values(r.Conn.Do(HGETALL, key))
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, repo.ErrNotFound
	}

	if err := redis.ScanStruct(v, &mod); err != nil {
		return nil, err
	}

	return &mod, nil
}

// multi is used internally to begin an atomic sequence
func (r *Repo) multi() error {

	v, err := redis.String(r.Conn.Do(MULTI))
	if err != nil {
		return err
	}
	if v != "OK" {
		return errors.New("multi failed")
	}

	return nil
}

func (r *Repo) exec() error {

	_, err := r.Conn.Do(EXEC)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) discard() error {
	v, err := redis.String(r.Conn.Do(DISCARD))
	if err != nil {
		return err
	}
	if v != "OK" {
		return errors.New("Discard failed")
	}

	return nil
}

///////////
///////////

// OpenConn .
func (r *Repo) OpenConn() error {
	logrus.Debug("redisrepo.OpenConn()")
	r.Conn = r.Pool.Get()
	if _, err := r.Conn.Do(AUTH, RedisPassword); err != nil {
		return err
	}
	return nil
}

// CloseConn .
func (r *Repo) CloseConn() error {
	logrus.Debug("redisrepo.CloseConn()")
	r.Conn.Close()
	return nil
}

// Close implements HatRepo.Close()
func (r *Repo) Close() error {
	logrus.Debug("redisrepo.Close()")
	r.Pool.Close()
	return nil
}
