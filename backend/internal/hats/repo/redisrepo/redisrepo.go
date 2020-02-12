package redisrepo

import (
	"backend/internal/hats/repo"
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

//Repo implements repo.HatRepo
type Repo struct {
	conn redis.Conn
}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

// NewRepo returns a pointer to a new instance of Repo
func NewRepo(conn redis.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

//FindAll queries all records
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {
	return hats, nil
}

//Save performs an upsert
func (r *Repo) Save(hm *repo.HatMod) error {

	id := fmt.Sprintf("hat:%s", hm.ID)

	if _, err := r.conn.Do("HMSET", redis.Args{}.Add(id).AddFlat(*hm)...); err != nil {
		return err
	}
	return nil
}

//BeginTx implements HatRepo.BeginTxn()
func (r *Repo) BeginTxn(ctx context.Context) error {
	logrus.Debug("redis.BeginTxn()")
	return nil
}

//Rollback implements HatRepo.Rollback()
func (r *Repo) Rollback() error {
	logrus.Debug("redis.Rollback()")
	return nil
}

//Commit implements HatRepo.Commit()
func (r *Repo) Commit() error {
	logrus.Debug("redis.Commit()")
	return nil
}
