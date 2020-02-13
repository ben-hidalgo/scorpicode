package redisrepo

import (
	"backend/internal/hats/repo"
	"context"
	"fmt"

	"github.com/rs/xid"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// HAT is the prefix of the hashmap UIDs
const HAT = "hat"

// HATS is the sortable range of hat UIDs
const HATS = "hats"

// returns the id and the key, creating a new one if empty
// e.g. id -> "123" and key -> "hat:123"
func idkey(inputID string) (id string, key string) {
	if inputID == "" {
		id = xid.New().String()
	} else {
		id = inputID
	}
	return id, fmt.Sprintf("%s:%s", HAT, id)
}

// FindAll queries all records
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {

	// values will be an array of strings
	values, err := redis.Values(r.conn.Do(SORT, HATS))
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
func (r *Repo) Save(hm repo.HatMod) (string, error) { //TODO: should we return a UUID and populate the ID here (rather than in the service)???

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
		return "", err
	}
	// TODO: wrap with if not EXISTS hat:123
	if _, err := r.conn.Do(LPUSH, HATS, hm.ID); err != nil {
		return "", err
	}

	return id, nil
}

// Delete deletes the record if version matches; throws NotFound, VersionMismatch
func (r *Repo) Delete(id string, version int) error {
	// TODO: implement
	return nil
}

// Exists returns true if the record exists
func (r *Repo) Exists(id string) (bool, error) {
	// TODO: implement
	ok := false
	return ok, nil
}

// Find one; returns NotFound
func (r *Repo) Find(id string) (*repo.HatMod, error) {
	return nil, nil
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
