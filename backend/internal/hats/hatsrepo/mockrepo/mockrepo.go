package mockrepo

import (
	"backend/internal/hats/hatsrepo"
	"context"
)

//FuncRepo mock implementing HatsRepo with all methods injectable
type FuncRepo struct {
	CreateHatF          func(context.Context, *hatsrepo.Hat) error
	CreateMakeHatsCmdF  func(context.Context, *hatsrepo.MakeHatsCmd) error
	DeleteMakeHatsCmdF  func(context.Context, *hatsrepo.MakeHatsCmd) error
	FindAllMakeHatsCmdF func(context.Context) ([]*hatsrepo.MakeHatsCmd, error)
	FindOneMakeHatsCmdF func(context.Context, string) (*hatsrepo.MakeHatsCmd, error)
	VisitTxnF           func(context.Context, func() error) error
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*FuncRepo)(nil)

// NewRepo .
func NewRepo() *FuncRepo {
	return &FuncRepo{}
}

// CreateHat calls the injected function
func (r *FuncRepo) CreateHat(ctx context.Context, h *hatsrepo.Hat) error {
	return r.CreateHatF(ctx, h)
}

// CreateMakeHatsCmd calls the injected function
func (r *FuncRepo) CreateMakeHatsCmd(ctx context.Context, mhc *hatsrepo.MakeHatsCmd) error {
	return r.CreateMakeHatsCmdF(ctx, mhc)
}

// DeleteMakeHatsCmd calls the injected function
func (r *FuncRepo) DeleteMakeHatsCmd(ctx context.Context, mhc *hatsrepo.MakeHatsCmd) error {
	return r.DeleteMakeHatsCmdF(ctx, mhc)
}

// FindOneMakeHatsCmd calls the injected function
func (r *FuncRepo) FindOneMakeHatsCmd(ctx context.Context, id string) (*hatsrepo.MakeHatsCmd, error) {
	return r.FindOneMakeHatsCmdF(ctx, id)
}

// FindAllMakeHatsCmd calls the injected function
func (r *FuncRepo) FindAllMakeHatsCmd(ctx context.Context) ([]*hatsrepo.MakeHatsCmd, error) {
	return r.FindAllMakeHatsCmdF(ctx)
}

// VisitTxn .
func (r *FuncRepo) VisitTxn(ctx context.Context, tf func() error) error {
	return r.VisitTxnF(ctx, tf)
}
