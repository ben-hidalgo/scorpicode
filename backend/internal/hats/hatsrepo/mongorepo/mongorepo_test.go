package mongorepo_test

import (
	"backend/internal/hats/hatsrepo"
	"backend/internal/hats/hatsrepo/mongorepo"
	"testing"
	"time"
)

const (
	NOT_NIL   = "not nil"
	NOT_EMPTY = "not empty"
	GOT       = "got '%v' %s '%v'"
	WANTED    = "but wanted"
)

func TestSaveHat(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	r := mongorepo.NewRepo()

	color := "RED"
	size := "06000"
	style := "DERBY"

	hat := &hatsrepo.Hat{
		Color: color,
		Size:  size,
		Style: style,
	}

	err := r.CreateHat(hat)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if hat.GetID() == nil {
		t.Fatalf(GOT, hat.GetID(), WANTED, NOT_NIL)
	}
	if hat.CreatedAt == (time.Time{}) {
		t.Fatalf(GOT, hat.CreatedAt, WANTED, NOT_EMPTY)
	}
	if hat.UpdatedAt == (time.Time{}) {
		t.Fatalf(GOT, hat.UpdatedAt, WANTED, NOT_EMPTY)
	}
	if hat.Version != 0 {
		t.Fatalf(GOT, hat.Version, WANTED, 0)
	}
	if hat.Color != color {
		t.Fatalf(GOT, hat.Color, WANTED, color)
	}
	if hat.Size != size {
		t.Fatalf(GOT, hat.Size, WANTED, size)
	}
	if hat.Style != style {
		t.Fatalf(GOT, hat.Style, WANTED, style)
	}

}
