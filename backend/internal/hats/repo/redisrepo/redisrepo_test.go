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

	if _, err := conn.Do(redisrepo.AUTH, redisrepo.RedisPassword); err != nil {
		t.Fatalf("auth failed password=%s err=%#v", redisrepo.RedisPassword, err)
	}

	reply, err := conn.Do(redisrepo.PING)
	pingReply, err := redis.String(reply, err)
	if pingReply != redisrepo.PONG {
		t.Fatalf("ping failed err=%#v", err)
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

	// hats, err := hr.FindAll(repo.Limit(10), repo.Offset(0))

}
