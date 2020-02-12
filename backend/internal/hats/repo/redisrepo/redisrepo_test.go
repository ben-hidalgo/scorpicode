package redisrepo_test

import (
	"backend/internal/hats/config"
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"context"
	"log"
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestCountMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	v, ok := os.LookupEnv("REDIS_ADDRESS")
	if !ok {
		t.Fail()
	}
	log.Printf(v)

	conn, err := redis.Dial("tcp", config.RedisAddress)
	if err != nil {
		t.Fatalf("failed to dial redis connection at %s err=%#v", config.RedisAddress, err)
	}
	defer conn.Close()

	if _, err := conn.Do("AUTH", config.RedisPassword); err != nil {
		t.Fatalf("auth failed password=%s err=%#v", config.RedisPassword, err)
	}

	reply, err := conn.Do("PING")
	if err != nil {
		t.Fatalf("failed to dial redis connection at %s err=%#v", config.RedisAddress, err)
	}
	_ = reply

	hr := redisrepo.NewRepo(conn)

	ctx := context.Background()

	hr.BeginTxn(ctx)

	id := "123"
	inches := int32(10)
	name := "cap"
	color := "blue"

	mod := &repo.HatMod{
		ID:     id,
		Color:  color,
		Name:   name,
		Inches: inches,
	}

	err = hr.Save(mod)
	if err != nil {
		t.Fatalf("save failed err=%#v", err)
	}

}
