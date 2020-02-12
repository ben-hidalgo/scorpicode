package redisrepo

import (
	"backend/internal/hats/repo"
	"backend/pkg/envconfig"

	"github.com/gomodule/redigo/redis"
)

//Repo implements repo.HatRepo
type Repo struct {
	conn redis.Conn
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

// NewConn initializes the redis connection
func NewConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", RedisAddress)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewRepo returns a pointer to a new instance of Repo
func NewRepo(conn redis.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

const (
	// AUTH .
	AUTH = "AUTH"
	// HMSET .
	HMSET = "HMSET"
	// PING .
	PING = "PING"
	// PONG .
	PONG = "PONG"
)
