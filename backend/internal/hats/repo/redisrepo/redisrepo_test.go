package redisrepo_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"context"
	"testing"
)

func TestSaveFind(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	hr := redisrepo.NewRepo()

	ctx := context.Background()

	hr.BeginTxn(ctx)

	inches := int32(10)
	name := "cap"
	color := "blue"

	mod := &repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: inches,
	}

	id, err := hr.Save(*mod)
	if err != nil {
		t.Fatalf("Save failed err=%#v", err)
	}
	if id == "" {
		t.Fatalf("id empty")
	}

	hats, err := hr.FindAll(repo.Limit(10), repo.Offset(0))
	if err != nil {
		t.Fatalf("FindAll failed err=%#v", err)
	}

	if len(hats) != 1 {
		t.Fatalf("unexpected len(hats)=%d", len(hats))
	}

}
