package mongorepo_test

import (
	"backend/internal/hats/hatsrepo"
	"backend/internal/hats/hatsrepo/mongorepo"
	"testing"
)

func TestPlaceholderExample(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	r := mongorepo.NewRepo()

	hat := &hatsrepo.Hat{
		Color:   "RED",
		Size:    "06000",
		Style:   "DERBY",
		Version: 1,
	}

	err := r.SaveHat(hat)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	// t.Fatalf("%#v", hat)

}
