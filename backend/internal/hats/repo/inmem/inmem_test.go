package inmem_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"backend/rpc/hatspb"
	"errors"
	"testing"
)

const (
	DefaultColor  = "red"
	DefaultStyle  = hatspb.Style_BASEBALL
	DefaultInches = int32(10)

	NOT_NIL = "not nil"
	GOT     = "got '%v' %s '%v'"
	WANTED  = "but wanted"
)

func start() (*inmem.Repo, *repo.HatMod) {
	hm := &repo.HatMod{
		Color:   DefaultColor,
		Style:   DefaultStyle.String(),
		Inches:  DefaultInches,
		Version: 0,
	}
	hr := inmem.NewRepo()
	for k := range hr.Storage {
		delete(hr.Storage, k)
	}
	return hr, hm
}

func TestNotExists(t *testing.T) {

	hr, _ := start()

	exists, err := hr.Exists("123")
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if exists != false {
		t.Fatalf(GOT, exists, WANTED, false)
	}
}

func TestExists(t *testing.T) {

	hr, hm := start()

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

	hr, _ := start()

	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if len(hats) != 0 {
		t.Fatalf(GOT, len(hats), WANTED, 0)
	}
}

func TestFindAllOne(t *testing.T) {

	hr, hm := start()

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
	if hat.Style != DefaultStyle.String() {
		t.Fatalf(GOT, hat.Style, WANTED, DefaultStyle.String())
	}
	if hat.Color != DefaultColor {
		t.Fatalf(GOT, hat.Color, WANTED, DefaultColor)
	}
}

func TestDeleteNotFound(t *testing.T) {

	hr, _ := start()

	err := hr.Delete("123", 0)
	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if !errors.Is(err, repo.ErrNotFound) {
		t.Fatalf(GOT, err, WANTED, repo.ErrNotFound)
	}
}

func TestDeleteFound(t *testing.T) {

	hr, hm := start()

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

	hr, hm := start()

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

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if mod.ID == "" {
		t.Fatalf(GOT, mod.ID, WANTED, "not nil")
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
	if hat.Style != DefaultStyle.String() {
		t.Fatalf(GOT, hat.Style, WANTED, DefaultStyle.String())
	}
	if hat.Color != DefaultColor {
		t.Fatalf(GOT, hat.Color, WANTED, DefaultColor)
	}
}

func TestSaveUpdate(t *testing.T) {

	hr, hm := start()

	mod, err := hr.Save(*hm)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	newColor := "blue"
	newStyle := hatspb.Style_BOWLER.String()
	newInches := int32(12)

	mod.Color = newColor
	mod.Style = newStyle
	mod.Inches = newInches

	hat, err := hr.Save(*mod)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if hat == nil {
		t.Fatalf(GOT, hat, WANTED, NOT_NIL)
	}
	if hat.ID != mod.ID {
		t.Fatalf(GOT, hat.ID, WANTED, mod.ID)
	}
	if hat.Inches != newInches {
		t.Fatalf(GOT, hat.Inches, WANTED, newInches)
	}
	if hat.Style != newStyle {
		t.Fatalf(GOT, hat.Style, WANTED, newStyle)
	}
	if hat.Color != newColor {
		t.Fatalf(GOT, hat.Color, WANTED, newColor)
	}
	if hat.Version != 2 {
		t.Fatalf(GOT, hat.Version, WANTED, 2)
	}
}

func TestSaveVersionMismatch(t *testing.T) {

	hr, hm := start()

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
