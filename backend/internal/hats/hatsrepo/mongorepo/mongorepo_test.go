package mongorepo_test

import (
	"backend/internal/hats/hatsrepo/mongorepo"
	"testing"
)

func TestPlaceholderExample(t *testing.T) {

	err := mongorepo.Init()
	if err != nil {
		t.Fatalf("%s", err)
	}
}
