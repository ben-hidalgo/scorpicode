package hatsrepo_test

import (
	"backend/internal/hats/hatsrepo"
	"testing"

	"github.com/Kamva/mgm/v2"
)

//FuncRepo mock implementing HatsRepo with all methods injectable
type FuncRepo struct {
	FuncSave func(b *hatsrepo.Book) (*mgm.IDField, error)
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*FuncRepo)(nil)

// Save calls the injected function
func (r *FuncRepo) Save(b *hatsrepo.Book) (*mgm.IDField, error) {
	return r.FuncSave(b)
}

// the service tests will use the FuncRepo mock
func TestPlaceholderExample(t *testing.T) {

	fr := &FuncRepo{}

	// inject mock function
	fr.FuncSave = func(b *hatsrepo.Book) (*mgm.IDField, error) {
		return nil, nil
	}

	id, err := fr.Save(&hatsrepo.Book{})

	if id != nil {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
