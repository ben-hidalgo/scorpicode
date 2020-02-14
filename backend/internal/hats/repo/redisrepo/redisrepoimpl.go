package redisrepo

import (
	"backend/internal/hats/repo"
	"context"
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

	// values will be an array of strings
	values, err := redis.Values(r.conn.Do(SORT, SetName))
	if err != nil {
		return
	}

	// TODO: v is a string, better naming and/or conversion
	for _, v := range values {
		// logrus.Printf("v=%s", v)

		_, key := idkey(fmt.Sprintf("%s", v))
		all, err := redis.Values(r.conn.Do(HGETALL, key))
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
func (r *Repo) Save(hm repo.HatMod) (*repo.HatMod, error) { //TODO: should we return a UUID and populate the ID here (rather than in the service)???

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

	if _, err := r.conn.Do(HMSET, redis.Args{}.Add(key).AddFlat(mod)...); err != nil {
		return nil, err
	}

	// set add hats <id>
	if _, err := r.conn.Do(SADD, SetName, mod.ID); err != nil {
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

	// delete the id from the set (hats)

	// value is an integer: 1 or 0; the number of values deleted
	vs, err := redis.Int(r.conn.Do(SREM, SetName, id))
	if err != nil {
		return err
	}
	if vs == 0 {
		return repo.ErrNotFound
	}

	// delete the key hat:<id>
	_, key := idkey(id)
	vk, err := redis.Int(r.conn.Do(DEL, key))
	if err != nil {
		return err
	}
	if vk == 0 {
		return repo.ErrNotFound
	}

	return nil
}

// Exists returns true if the record exists
func (r *Repo) Exists(id string) (bool, error) {

	// set is member
	// value is an integer: 1 or 0
	v, err := redis.Int(r.conn.Do(SISMEMBER, SetName, id))
	if err != nil {
		return false, err
	}

	return v == 1, nil
}

// Find one; returns NotFound
func (r *Repo) Find(id string) (*repo.HatMod, error) {

	var mod repo.HatMod

	_, key := idkey(id)

	v, err := redis.Values(r.conn.Do(HGETALL, key))
	if err != nil {
		return nil, err
	}

	if err := redis.ScanStruct(v, &mod); err != nil {
		return nil, err
	}

	return &mod, nil
}

///////////
///////////

// BeginTxn implements HatRepo.BeginTxn()
func (r *Repo) BeginTxn(ctx context.Context) error {
	logrus.Debug("BeginTxn()")
	return nil
}

// Rollback implements HatRepo.Rollback()
func (r *Repo) Rollback() error {
	logrus.Debug("Rollback()")
	return nil
}

// Commit implements HatRepo.Commit()
func (r *Repo) Commit() error {
	logrus.Debug("Commit()")
	return nil
}

// Close implements HatRepo.Close()
func (r *Repo) Close() error {
	logrus.Debug("Close()")
	r.conn.Close()
	return nil
}
