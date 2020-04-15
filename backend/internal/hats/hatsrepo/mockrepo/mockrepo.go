package mockrepo

import (
	"backend/internal/hats/hatsrepo"
)

//FuncRepo mock implementing HatsRepo with all methods injectable
type FuncRepo struct {
	CreateHatF          func(h *hatsrepo.Hat) error
	CreateMakeHatsCmdF  func(mhc *hatsrepo.MakeHatsCmd) error
	DeleteMakeHatsCmdF  func(id string, version int32) error
	FindAllMakeHatsCmdF func() ([]*hatsrepo.MakeHatsCmd, error)
	FindOneMakeHatsCmdF func(id string) (*hatsrepo.MakeHatsCmd, error)
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

// DeleteMakeHatsCmd calls the injected function
func (r *FuncRepo) DeleteMakeHatsCmd(id string, version int32) error {
	return r.DeleteMakeHatsCmdF(id, version)
}

// FindOneMakeHatsCmd calls the injected function
func (r *FuncRepo) FindOneMakeHatsCmd(id string) (*hatsrepo.MakeHatsCmd, error) {
	return r.FindOneMakeHatsCmdF(id)
}

// FindAllMakeHatsCmd calls the injected function
func (r *FuncRepo) FindAllMakeHatsCmd() ([]*hatsrepo.MakeHatsCmd, error) {
	return r.FindAllMakeHatsCmdF()
}
