package redisrepo

import (
	"backend/internal/hats/repo"
	"backend/pkg/envconfig"

	"github.com/sirupsen/logrus"

	"github.com/gomodule/redigo/redis"
)

//Repo implements repo.HatRepo
type Repo struct {
	Conn redis.Conn
}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

// RedisAddress .
var RedisAddress = ""

// RedisPassword .
var RedisPassword = ""

func init() {
	envconfig.SetString("REDIS_ADDRESS", &RedisAddress)
	envconfig.SetString("REDIS_PASSWORD", &RedisPassword)
}

// NewRepo returns a pointer to a new instance of Repo
// will panic on connection errors
func NewRepo() *Repo {
	conn, err := redis.Dial("tcp", RedisAddress)
	if err != nil {
		logrus.Panicf("failed to dial redis connection at %s err=%#v", RedisAddress, err)
	}

	if _, err := conn.Do(AUTH, RedisPassword); err != nil {
		logrus.Panicf("auth failed err=%#v", err)
	}

	reply, err := redis.String(conn.Do(PING))
	if reply != PONG {
		logrus.Panicf("unexpected reply to PING err=%#v", err)
	}

	return &Repo{
		Conn: conn,
	}
}

const (
	// AUTH authorize
	AUTH = "AUTH"

	// PING ping
	PING = "PING"

	// PONG ping response
	PONG = "PONG"

	// MULTI begin txn
	MULTI = "MULTI"

	// DISCARD rollback
	DISCARD = "DISCARD"

	// EXEC commit
	EXEC = "EXEC"

	// FLUSHDB delete all keys in db
	FLUSHDB = "FLUSHDB"

	// HMSET hashmap set
	HMSET = "HMSET"

	// LPUSH list push
	LPUSH = "LPUSH"

	// DEL delete key
	DEL = "DEL"

	// SORT sort
	SORT = "SORT"

	// HGETALL hash get all
	HGETALL = "HGETALL"

	// SADD set add
	SADD = "SADD"

	// SREM set remove
	SREM = "SREM"

	// SISMEMBER set is member
	SISMEMBER = "SISMEMBER"

	// SMEMBERS set members
	SMEMBERS = "SMEMBERS"
)
