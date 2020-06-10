package hatdao_test

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/mongodb"
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultColor   = "RED"
	DefaultStyle   = "FEDORA"
	DefaultSize    = "06000"
	DefaultOrderID = "5e8e20bbe6b38b8cb0870808"
	DefaultSubject = "google-oauth2|204673116516641832842"
)

const (
	NOT_NIL   = "not nil"
	NOT_EMPTY = "not empty"
	GOT       = "got '%v' %s '%v'"
	WANTED    = "but wanted"
)

func TestCreate(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	// connect mongo
	mongoClient, err := mongodb.Client()
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	defer mongoClient.Disconnect(context.Background())

	dao := hatdao.New(mongoClient)

	id, _ := primitive.ObjectIDFromHex(DefaultOrderID)

	hat := &hatdao.Hat{
		Color:     DefaultColor,
		Style:     DefaultStyle,
		Size:      DefaultSize,
		Ordinal:   1, // one indexed
		OrderID:   id,
		CreatedBy: DefaultSubject,
	}
	err = dao.Create(context.Background(), hat)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if hat.GetID() == nil {
		t.Fatalf(GOT, hat.GetID(), WANTED, NOT_NIL)
	}
	if hat.CreatedAt == (time.Time{}) {
		t.Fatalf(GOT, hat.CreatedAt, WANTED, NOT_EMPTY)
	}
	if hat.UpdatedAt == (time.Time{}) {
		t.Fatalf(GOT, hat.UpdatedAt, WANTED, NOT_EMPTY)
	}
	if hat.Version != 1 {
		t.Fatalf(GOT, hat.Version, WANTED, 1)
	}
	if hat.Size != DefaultSize {
		t.Fatalf(GOT, hat.Size, WANTED, DefaultSize)
	}
	if hat.Style != DefaultStyle {
		t.Fatalf(GOT, hat.Style, WANTED, DefaultStyle)
	}
	if hat.Color != DefaultColor {
		t.Fatalf(GOT, hat.Color, WANTED, DefaultColor)
	}

}
