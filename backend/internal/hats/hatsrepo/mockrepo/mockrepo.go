package mockrepo

import (
	"backend/internal/hats/hatsrepo"
)

//FuncRepo mock implementing HatsRepo with all methods injectable
type FuncRepo struct {
	CreateHatF         func(h *hatsrepo.Hat) error
	CreateMakeHatsCmdF func(mhc *hatsrepo.MakeHatsCmd) error
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*FuncRepo)(nil)

// NewRepo .
func NewRepo() *FuncRepo {
	return &FuncRepo{}
}

// CreateHat calls the injected function
func (r *FuncRepo) CreateHat(h *hatsrepo.Hat) error {
	return r.CreateHatF(h)
}

// CreateMakeHatsCmd calls the injected function
func (r *FuncRepo) CreateMakeHatsCmd(mhc *hatsrepo.MakeHatsCmd) error {
	return r.CreateMakeHatsCmdF(mhc)
}
