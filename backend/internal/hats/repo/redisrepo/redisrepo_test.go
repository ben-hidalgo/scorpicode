package redisrepo_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"context"
	"testing"
)

func TestCountMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	hr := redisrepo.NewRepo()

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

	err := hr.Save(mod)
	if err != nil {
		t.Fatalf("save failed err=%#v", err)
	}

	// hats, err := hr.FindAll(repo.Limit(10), repo.Offset(0))

}
