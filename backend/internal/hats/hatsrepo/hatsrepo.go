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

// HatsRepo .
type HatsRepo interface {
	Save(b *Book) (*mgm.IDField, error)
}

// NewBook .
func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}
