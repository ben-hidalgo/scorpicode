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

	// instantiate the hatdao
	dao := hatdao.New(mongoClient)

	// Not Found
	hnf, err := dao.Find(context.Background(), "123")
	if err != nil {
		// this enforces that not found does not return an error
		t.Fatalf(GOT, err, WANTED, nil)
		// this is a design decision to avoid having to check for err == NotFoundError
		// and differentiates actual failures from successfully retrieving nothing
	}
	// if the entity is not found, the response will be nil
	if hnf != nil {
		t.Fatalf(GOT, hnf, WANTED, nil)
	}

	// Create
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

	// Found
	h1, err := dao.Find(context.Background(), hat.ID.Hex())
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if h1.GetID() == nil {
		t.Fatalf(GOT, h1.GetID(), WANTED, NOT_NIL)
	}
	if h1.CreatedAt == (time.Time{}) {
		t.Fatalf(GOT, h1.CreatedAt, WANTED, NOT_EMPTY)
	}
	if h1.UpdatedAt == (time.Time{}) {
		t.Fatalf(GOT, h1.UpdatedAt, WANTED, NOT_EMPTY)
	}
	if h1.Version != 1 {
		t.Fatalf(GOT, h1.Version, WANTED, 1)
	}
	if h1.Size != DefaultSize {
		t.Fatalf(GOT, h1.Size, WANTED, DefaultSize)
	}
	if h1.Style != DefaultStyle {
		t.Fatalf(GOT, h1.Style, WANTED, DefaultStyle)
	}
	if h1.Color != DefaultColor {
		t.Fatalf(GOT, h1.Color, WANTED, DefaultColor)
	}

}
