package inmem_test

import (
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"testing"
)

func TestRepo(t *testing.T) {

	hr := inmem.NewRepo()

	id := "123"
	color := "red"
	name := "cap"
	inches := int32(10)

	mod := &repo.HatMod{
		ID:     id,
		Color:  color,
		Name:   name,
		Inches: inches,
	}

	found, err := hr.Exists(id)
	if err != nil {
		t.Fatalf("err should be nil %#v", err)
	}
	if found {
		t.Fatalf("should not be found")
	}

	// find when empty returns length zero
	hats, err := hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf("save failed err=%s", err)
	}
	if len(hats) != 0 {
		t.Fatalf("length should be zero but was %d", len(hats))
	}

	// save the first one
	err = hr.Save(mod)
	if err != nil {
		t.Fatalf("err should be nil %#v", err)
	}

	found, err = hr.Exists(id)
	if err != nil {
		t.Fatalf("err should be nil %#v", err)
	}
	if !found {
		t.Fatalf("should be found")
	}

	// find the one we just saved
	hats, err = hr.FindAll(10, 0)
	if err != nil {
		t.Fatalf("save failed err=%s", err)
	}
	if len(hats) != 1 {
		t.Fatalf("length should be zero but was %d", len(hats))
	}

	// validate the hat is returned correctly
	hat := hats[0]

	if hat == nil {
		t.Fatalf("hat should not be nil")
	}
	if hat.Inches != inches {
		t.Fatalf("inches must match %d != %d", hat.Inches, inches)
	}
	if hat.Name != name {
		t.Fatalf("name must match %s != %s", hat.Name, name)
	}
	if hat.Color != color {
		t.Fatalf("color must match %s != %s", hat.Color, color)
	}

}
