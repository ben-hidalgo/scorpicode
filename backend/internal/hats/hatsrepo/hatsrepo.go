package hatsrepo

import (
	"context"
	"errors"

	"github.com/Kamva/mgm/v2"
)

// ErrVersionMismatch is for optimistic locking version mismatch on update and delete
var ErrVersionMismatch = errors.New("err.version.mismatch")

// Hat .
type Hat struct {
	// DefaultModel includes: _id,created_at and updated_at
	mgm.DefaultModel `bson:",inline"`
	Size             string `json:"size"          bson:"size"`
	Color            string `json:"color"         bson:"color"`
	Style            string `json:"style"         bson:"style"`
	MakeHatsCmdID    string `json:"makeHatsCmdId" bson:"makeHatsCmdId"`
	Version          int32  `json:"version"       bson:"version"`
}

// MakeHatsCmd .
type MakeHatsCmd struct {
	// DefaultModel includes: _id,created_at and updated_at
	mgm.DefaultModel `bson:",inline"`
	Size             string `json:"size"     bson:"size"`
	Color            string `json:"color"    bson:"color"`
	Style            string `json:"style"    bson:"style"`
	Quantity         int32  `json:"quantity" bson:"quantity"`
	Notes            string `json:"notes"    bson:"notes"`
	Version          int32  `json:"version"  bson:"version"`
}

// HatsRepo .
type HatsRepo interface {
	// TODO: all funcs should accept Context
	CreateHat(h *Hat) error
	CreateMakeHatsCmd(mhc *MakeHatsCmd) error
	DeleteMakeHatsCmd(id string, version int32) error
	FindAllMakeHatsCmd() ([]*MakeHatsCmd, error)
}

// used to store the Repo in Context
type key int

// RepoKey is the key for the repo in context; public for mock injection
const RepoKey key = 0

// FromContext returns the repo and panics if not found
func FromContext(ctx context.Context) HatsRepo {

	switch v := ctx.Value(RepoKey).(type) {
	case HatsRepo:
		return v
	default:
		panic("FromContext() no value found")
	}
}