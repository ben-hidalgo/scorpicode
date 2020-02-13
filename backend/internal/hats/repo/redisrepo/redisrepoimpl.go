package redisrepo

import (
	"backend/internal/hats/repo"
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// HAT is the prefix of the hashmap UIDs
const HAT = "hat"

// HATS is the sortable range of hat UIDs
const HATS = "hats"

func uid(hm *repo.HatMod) string {
	return suid(hm.ID)
}
func suid(id string) string {
	return fmt.Sprintf("%s:%s", HAT, id)
}

// FindAll queries all records
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {

	// values will be an array of strings
	values, err := redis.Values(r.conn.Do(SORT, HATS))
	if err != nil {
		return
	}

	for _, v := range values {
		logrus.Printf("v=%s", v)

		all, err := redis.Values(r.conn.Do(HGETALL, suid(fmt.Sprintf("%s", v))))
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
func (r *Repo) Save(hm *repo.HatMod) error {

	uid := uid(hm)

	if _, err := r.conn.Do(HMSET, redis.Args{}.Add(uid).AddFlat(*hm)...); err != nil {
		return err
	}
	if _, err := r.conn.Do(LPUSH, HATS, hm.ID); err != nil {
		return err
	}

	return nil
}

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
