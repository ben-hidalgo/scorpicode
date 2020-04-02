package mongorepo_test

import (
	"backend/internal/hats/hatsrepo"
	"backend/internal/hats/hatsrepo/mongorepo"
	"testing"

	"github.com/Kamva/mgm/v2"
)

func TestPlaceholderExample(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	err := mongorepo.Init()
	if err != nil {
		t.Fatalf("%s", err)
	}

	book := hatsrepo.NewBook("my book", 5)

	// Make sure pass the model by reference.
	err = mgm.Coll(book).Create(book)
	if err != nil {
		t.Fatalf("%s", err)
	}

}
