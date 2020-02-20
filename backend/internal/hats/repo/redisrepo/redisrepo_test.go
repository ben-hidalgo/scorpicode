package redisrepo_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/redisrepo"
	"errors"
	"testing"
)

const (
	DefaultColor  = "red"
	DefaultName   = "cap"
	DefaultInches = int32(10)

	NOT_NIL = "not nil"
	GOT     = "got '%v' %s '%v'"
	WANTED  = "but wanted"
)

func start(t *testing.T) (*redisrepo.Repo, *repo.HatMod) {
	hm := &repo.HatMod{
		Color:   DefaultColor,
		Name:    DefaultName,
		Inches:  DefaultInches,
		Version: 0,
	}
	hr := redisrepo.NewRepo()
	if err := hr.OpenConn(); err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if _, err := hr.Conn.Do(redisrepo.FLUSHDB); err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	return hr, hm
}

func TestNotExists(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, _ := start(t)

	exists, err := hr.Exists("123")
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if exists != false {
		t.Fatalf(GOT, exists, WANTED, false)
	}
}

func TestExists(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if exists != true {
		t.Fatalf(GOT, exists, WANTED, true)
	}
}

func TestFindAllEmpty(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, _ := start(t)

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if len(hats) != 0 {
		t.Fatalf(GOT, len(hats), WANTED, 0)
	}
}

func TestFindAllOne(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if len(hats) != 1 {
		t.Fatalf(GOT, len(hats), WANTED, 1)
	}

	// validate the correct hat is returned
	hat := hats[0]

	if hat == nil {
		t.Fatalf(GOT, hat, WANTED, NOT_NIL)
	}
	if hat.ID != mod.ID {
		t.Fatalf(GOT, hat.ID, WANTED, mod.ID)
	}
	if hat.Inches != DefaultInches {
		t.Fatalf(GOT, hat.Inches, WANTED, DefaultInches)
	}
	if hat.Name != DefaultName {
		t.Fatalf(GOT, hat.Name, WANTED, DefaultName)
	}
	if hat.Color != DefaultColor {
		t.Fatalf(GOT, hat.Color, WANTED, DefaultColor)
	}
}

func TestDeleteNotFound(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, _ := start(t)

	err := hr.Delete("123456", 0)
	if !errors.Is(err, repo.ErrNotFound) {
		t.Fatalf(GOT, err, WANTED, repo.ErrNotFound)
	}
}

func TestDeleteFound(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	err = hr.Delete(mod.ID, mod.Version)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if exists != false {
		t.Fatalf(GOT, exists, WANTED, false)
	}
}

func TestDeleteVersionMismatch(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	err = hr.Delete(mod.ID, -1)
	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if !errors.Is(err, repo.ErrVersionMismatch) {
		t.Fatalf(GOT, err, WANTED, repo.ErrVersionMismatch)
	}
}

func TestSaveInsert(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if len(mod.ID) == 0 {
		t.Fatalf(GOT, len(mod.ID), WANTED, "non-zero")
	}
	if mod.Version != 1 {
		t.Fatalf(GOT, mod.Version, WANTED, 1)
	}

	exists, err := hr.Exists(mod.ID)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if exists != true {
		t.Fatalf(GOT, exists, WANTED, true)
	}

	hat, err := hr.Find(mod.ID)

	if hat == nil {
		t.Fatalf(GOT, hat, WANTED, NOT_NIL)
	}
	if hat.ID != mod.ID {
		t.Fatalf(GOT, hat.ID, WANTED, mod.ID)
	}
	if hat.Inches != DefaultInches {
		t.Fatalf(GOT, hat.Inches, WANTED, DefaultInches)
	}
	if hat.Name != DefaultName {
		t.Fatalf(GOT, hat.Name, WANTED, DefaultName)
	}
	if hat.Color != DefaultColor {
		t.Fatalf(GOT, hat.Color, WANTED, DefaultColor)
	}
}

func TestSaveVersionMismatch(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	hr, hm := start(t)

	// Version is non-zero while ID is empty string
	hm.Version = 1

	_, err := hr.Save(*hm)
	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if !errors.Is(err, repo.ErrVersionMismatch) {
		t.Fatalf(GOT, err, WANTED, repo.ErrVersionMismatch)
	}
}
