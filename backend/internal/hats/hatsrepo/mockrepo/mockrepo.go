package mockrepo

import (
	"backend/internal/hats/hatsrepo"
)

//FuncRepo mock implementing HatsRepo with all methods injectable
type FuncRepo struct {
	SaveHatF func(h *hatsrepo.Hat) error
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*FuncRepo)(nil)

// SaveHat calls the injected function
func (r *FuncRepo) SaveHat(h *hatsrepo.Hat) error {
	return r.SaveHatF(h)
}
