package redisrepo

import (
	"backend/internal/hats/repo"
	"backend/pkg/envconfig"
	"time"

	"github.com/gomodule/redigo/redis"
)

//Repo implements repo.HatRepo
type Repo struct {
	Conn redis.Conn
	Pool *redis.Pool
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

	pool := newPool(RedisAddress)

	return &Repo{
		Pool: pool,
	}
}

func newPool(address string) *redis.Pool {

	return &redis.Pool{

		// TODO: env vars
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

const (
	// AUTH authorize
	AUTH = "AUTH"

	// PING ping
	PING = "PING"

	// PONG ping response
	PONG = "PONG"

	// MULTI begin atomic commands
	MULTI = "MULTI"

	// DISCARD discard multi
	DISCARD = "DISCARD"

	// EXEC exec multi
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
