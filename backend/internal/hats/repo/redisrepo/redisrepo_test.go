package redisrepo_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"context"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestCountMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conn, err := redisrepo.NewConn()
	if err != nil {
		t.Fatalf("failed to dial redis connection at %s err=%#v", redisrepo.RedisAddress, err)
	}
	defer conn.Close()

	if _, err := conn.Do("AUTH", redisrepo.RedisPassword); err != nil {
		t.Fatalf("auth failed password=%s err=%#v", redisrepo.RedisPassword, err)
	}

	reply, err := conn.Do("PING")
	if err != nil {
		t.Fatalf("ping failed err=%#v", err)
	}
	pingReply, err := redis.String(reply, err)
	if err != nil {
		t.Fatalf("ping reply string failed err=%#v", err)
	}
	if pingReply != "PONG" {
		t.Fatalf("expected %s but was %s", "PONG", pingReply)
	}

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
