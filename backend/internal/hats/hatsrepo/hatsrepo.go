package hatsrepo

import (
	"github.com/Kamva/mgm/v2"
)

// Book .
type Book struct {
	// DefaultModel includes: add _id,created_at and updated_at
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

// NewBook .
func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}

// HatsRepo .
type HatsRepo interface {

	// Save performs an upsert, returns the ID
	// Input parameter is not mutated
	// Assigns an ID if not provided (insert)
	// Increments Version
	// Returns NotFound if missing by ID
	// Returns VersionMismatch if version isn't equal
	Save(b *Book) (*mgm.IDField, error)
}

//RepoImpl implements HatsRepo
type RepoImpl struct {
}

// enforces the interface is implemented
var _ HatsRepo = (*RepoImpl)(nil)

// Save .
func (r *RepoImpl) Save(b *Book) (*mgm.IDField, error) {
	return nil, nil
}
