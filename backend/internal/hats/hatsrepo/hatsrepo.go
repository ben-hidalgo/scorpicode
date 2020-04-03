package hatsrepo

import (
	"github.com/Kamva/mgm/v2"
)

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
	SaveHat(h *Hat) error
}
