package mockrepo

import (
	"backend/internal/hats/hatsrepo"

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
