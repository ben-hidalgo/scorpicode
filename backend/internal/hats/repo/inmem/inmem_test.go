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
	EXPECTED  = "expected %v was %v"
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
		t.Fatalf(EXPECTED, nil, err)
	}
	if exists {
		t.Fatalf(EXPECTED, false, true)
	}

}

func TestExists(t *testing.T) {

	hr, hm := start()

	hm.ID = "123"

	err := hr.Save(hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}

	exists, err := hr.Exists("123")
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}
	if !exists {
		t.Fatalf(EXPECTED, true, false)
	}

}

func TestFindAllEmpty(t *testing.T) {

	hr, _ := start()

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}
	if len(hats) != 0 {
		t.Fatalf(EXPECTED, 0, len(hats))
	}
}

func TestFindAllOne(t *testing.T) {

	hr, hm := start()

	hm.ID = "123"

	err := hr.Save(hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}
	if len(hats) != 1 {
		t.Fatalf(EXPECTED, 1, len(hats))
	}

	// validate the correct hat is returned
	hat := hats[0]

	if hat == nil {
		t.Fatalf(EXPECTED, "!nil", nil)
	}
	if hat.ID != "123" {
		t.Fatalf(EXPECTED, "123", hat.ID)
	}
	if hat.Inches != expInches {
		t.Fatalf(EXPECTED, hat.Inches, expInches)
	}
	if hat.Name != expName {
		t.Fatalf(EXPECTED, hat.Name, expName)
	}
	if hat.Color != expColor {
		t.Fatalf(EXPECTED, hat.Color, expColor)
	}
}

func TestDeleteNotFound(t *testing.T) {

	hr, _ := start()

	err := hr.Delete("123", 0)
	if err == nil {
		t.Fatalf(EXPECTED, "!nil", nil)
	}
	if !errors.Is(err, repo.ErrNotFound) {
		t.Fatalf(EXPECTED, repo.ErrNotFound, err)
	}

}

func TestDeleteFound(t *testing.T) {

	hr, hm := start()

	id := "123"

	hm.ID = id

	err := hr.Save(hm)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}

	err = hr.Delete(id, 0)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}

	exists, err := hr.Exists(id)
	if err != nil {
		t.Fatalf(EXPECTED, nil, err)
	}
	if exists {
		t.Fatalf(EXPECTED, false, true)
	}

}
