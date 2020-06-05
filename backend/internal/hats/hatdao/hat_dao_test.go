package hatdao_test

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/mongodb"
	"context"
	"testing"
	"time"
)

const (
	DefaultColor    = "RED"
	DefaultStyle    = "FEDORA"
	DefaultSize     = "06000"
	DefaultQuantity = 10
	DefaultNotes    = "Lorem ipsum"
	DefaultHexID    = "5e8e20bbe6b38b8cb0870808"
	DefaultSubject  = "google-oauth2|204673116516641832842"
)

var DefaultCreatedAt = time.Date(2020, time.January, 0, 0, 0, 0, 0, time.UTC)
var DefaultUpdatedAt = time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)

/*
func startHat(mock *orderdaomock.Mock, rm *rabbitmock.Mock) (context.Context, *hatserver.Server, *hatspb.MakeHatsRequest) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, orderdao.Key, mock)

	ctx = context.WithValue(ctx, rabbit.Key, rm)

	ctx = context.WithValue(ctx, authnz.Key, &authnz.BearerToken{
		CC: &authnz.CustomClaims{
			Roles: []authnz.Role{
				authnz.HABERDASHER,
			},
			Claims: jwt.Claims{
				Subject: DefaultSubject,
			},
		},
	})

	hs := hatserver.NewServer()

	req := &hatspb.MakeHatsRequest{
		Size:     DefaultSize,
		Style:    DefaultStyle,
		Color:    DefaultColor,
		Quantity: DefaultQuantity,
		Notes:    DefaultNotes,
	}

	return ctx, hs, req
}
*/

func TestCreate(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	// connect mongo
	mongoClient, err := mongodb.Client()
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer mongoClient.Disconnect(context.Background())

	dao := hatdao.New(mongoClient)

	t.Logf("XXXXX %#v", mongoClient)
	t.Logf("XXXXX %#v", dao)

}
