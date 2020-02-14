package inmem_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"errors"
	"testing"
)

const (
	expColor  = "red"
	expName   = "cap"
	expInches = int32(10)
	EXPECTED  = "expected %v %s %v"
	BUT_WAS   = "but was"
	NOT_NIL   = "!nil"
)

func start() (*inmem.Repo, *repo.HatMod) {
	hm := &repo.HatMod{
		Color:   expColor,
		Name:    expName,
		Inches:  expInches,
		Version: 0,
	}
	return inmem.NewRepo(), hm
}

func TestNotExists(t *testing.T) {

	hr, _ := start()

	exists, err := hr.Exists("123")
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if exists {
		t.Fatalf(EXPECTED, false, BUT_WAS, true)
	}

}

func TestExists(t *testing.T) {

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if !exists {
		t.Fatalf(EXPECTED, true, BUT_WAS, false)
	}

}

func TestFindAllEmpty(t *testing.T) {

	hr, _ := start()

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if len(hats) != 0 {
		t.Fatalf(EXPECTED, 0, BUT_WAS, len(hats))
	}
}

func TestFindAllOne(t *testing.T) {

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if len(hats) != 1 {
		t.Fatalf(EXPECTED, 1, BUT_WAS, len(hats))
	}

	// validate the correct hat is returned
	hat := hats[0]

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

func TestDeleteNotFound(t *testing.T) {

	hr, _ := start()

	err := hr.Delete("123", 0)
	if err == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if !errors.Is(err, repo.ErrNotFound) {
		t.Fatalf(EXPECTED, repo.ErrNotFound, BUT_WAS, err)
	}

}

func TestDeleteFound(t *testing.T) {

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}

	err = hr.Delete(mod.ID, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if exists {
		t.Fatalf(EXPECTED, false, BUT_WAS, true)
	}

}

func TestSaveInsert(t *testing.T) {

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if !exists {
		t.Fatalf(EXPECTED, true, BUT_WAS, false)
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
