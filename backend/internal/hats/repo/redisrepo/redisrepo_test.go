package redisrepo_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"context"
	"testing"
)

const (
	expColor  = "red"
	expName   = "cap"
	expInches = int32(10)
	EXPECTED  = "expected %v %s %v"
	BUT_WAS   = "but was"
	NOT_NIL   = "!nil"
	NOT_EMPTY = "not empty"
)

func TestSaveInsert(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	hr := redisrepo.NewRepo()

	ctx := context.Background()

	hr.BeginTxn(ctx)

	inches := int32(10)
	name := "cap"
	color := "blue"

	mod, err := hr.Save(repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: inches,
	})
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if mod.ID == "" {
		t.Fatalf(EXPECTED, "", BUT_WAS, mod.ID)
	}

	hat, err := hr.Find(mod.ID)

	if hat == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if hat.ID != mod.ID {
		t.Fatalf(EXPECTED, mod.ID, BUT_WAS, hat.ID)
	}
	if hat.Inches != expInches {
		t.Fatalf(EXPECTED, hat.Inches, BUT_WAS, expInches)
	}
	if hat.Name != expName {
		t.Fatalf(EXPECTED, hat.Name, BUT_WAS, expName)
	}
	if hat.Color != expColor {
		t.Fatalf(EXPECTED, hat.Color, BUT_WAS, expColor)
	}
}
