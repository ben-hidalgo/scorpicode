package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/internal/hats/hatsrepo"
	"backend/internal/hats/hatsrepo/mockrepo"
	"backend/pkg/authnz"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"testing"
	"time"

	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultColor    = "RED"
	DefaultStyle    = hatspb.Style_FEDORA
	DefaultSize     = "06000"
	DefaultQuantity = 10
	DefaultNotes    = "Lorem ipsum"
	DefaultHexID    = "5e8e20bbe6b38b8cb0870808"
)

var DefaultCreatedAt = time.Date(2020, time.January, 0, 0, 0, 0, 0, time.UTC)
var DefaultUpdatedAt = time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)

func startHat(mr *mockrepo.FuncRepo) (context.Context, *hatserver.Server, *hatspb.MakeHatsRequest) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, hatsrepo.RepoKey, mr)

	ctx = context.WithValue(ctx, authnz.Key, &authnz.BearerToken{
		CC: &authnz.CustomClaims{
			Roles: []authnz.Role{
				authnz.Role("HABERDASHER"),
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

func TestHatSuccess(t *testing.T) {

	mr := &mockrepo.FuncRepo{
		CreateMakeHatsCmdF: func(mhc *hatsrepo.MakeHatsCmd) error {
			id, err := primitive.ObjectIDFromHex(DefaultHexID)
			if err != nil {
				t.Fatal(err)
			}
			mhc.SetID(id)
			mhc.CreatedAt = DefaultCreatedAt
			mhc.UpdatedAt = DefaultUpdatedAt
			mhc.Version = 1
			return nil
		},
	}

	ctx, hs, req := startHat(mr)

	res, err := hs.MakeHats(ctx, req)

	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if res == nil {
		t.Fatalf(GOT, res, WANTED, NOT_NIL)
	}
	if res.GetHat() == nil {
		t.Fatalf(GOT, res.GetHat(), WANTED, NOT_NIL)
	}
	if res.GetHat().GetId() != DefaultHexID {
		t.Fatalf(GOT, res.GetHat().GetId(), WANTED, DefaultHexID)
	}
	if res.GetHat().GetCreatedAt() != DefaultCreatedAt.Format(time.RFC3339) {
		t.Fatalf(GOT, res.GetHat().GetCreatedAt(), WANTED, DefaultCreatedAt.Format(time.RFC3339))
	}
	if res.GetHat().GetUpdatedAt() != DefaultUpdatedAt.Format(time.RFC3339) {
		t.Fatalf(GOT, res.GetHat().GetUpdatedAt(), WANTED, DefaultUpdatedAt.Format(time.RFC3339))
	}
	if res.GetHat().GetVersion() != 1 {
		t.Fatalf(GOT, res.GetHat().GetVersion(), WANTED, 1)
	}
	if res.GetHat().GetSize() != DefaultSize {
		t.Fatalf(GOT, res.GetHat().GetSize(), WANTED, DefaultSize)
	}
	if res.GetHat().GetStyle() != DefaultStyle {
		t.Fatalf(GOT, res.GetHat().GetStyle(), WANTED, DefaultStyle)
	}
	if res.GetHat().GetColor() != DefaultColor {
		t.Fatalf(GOT, res.GetHat().GetColor(), WANTED, DefaultColor)
	}

}

func TestMissingRole(t *testing.T) {

	ctx, hs, req := startHat(mockrepo.NewRepo())

	// overwrite the default bearer with no roles
	ctx = context.WithValue(ctx, authnz.Key, &authnz.BearerToken{
		CC: nil,
	})

	res, err := hs.MakeHats(ctx, req)

	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if res != nil {
		t.Fatalf(GOT, res, WANTED, nil)
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.PermissionDenied {
		t.Fatalf(GOT, e.Code(), WANTED, twirp.PermissionDenied)
	}
	if e.Msg() != string(hatserver.MakeHatsForbidden) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.MakeHatsForbidden)
	}

}

func TestSizeRequired(t *testing.T) {

	ctx, hs, req := startHat(mockrepo.NewRepo())

	req.Size = ""

	testRequired(t, ctx, hs, req, hatserver.HatSizeRequired)
}

func TestColorRequired(t *testing.T) {

	ctx, hs, req := startHat(mockrepo.NewRepo())

	req.Color = ""

	testRequired(t, ctx, hs, req, hatserver.HatColorRequired)
}

func TestColorDomain(t *testing.T) {

	ctx, hs, req := startHat(mockrepo.NewRepo())

	req.Color = "not a color"

	res, err := hs.MakeHats(ctx, req)

	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if res != nil {
		t.Fatalf(GOT, res, WANTED, nil)
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf(GOT, e.Code(), WANTED, twirp.InvalidArgument)
	}
	if e.Msg() != string(hatserver.HatColorDomain) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.HatColorDomain)
	}
}

func TestNameRequired(t *testing.T) {

	ctx, hs, req := startHat(mockrepo.NewRepo())

	req.Style = hatspb.Style_UNKNOWN_STYLE

	testRequired(t, ctx, hs, req, hatserver.HatStyleRequired)
}

func testRequired(t *testing.T, ctx context.Context, hs *hatserver.Server, req *hatspb.MakeHatsRequest, emsg util.ErrMsg) {

	res, err := hs.MakeHats(ctx, req)

	if err == nil {
		t.Fatalf(GOT, err, WANTED, NOT_NIL)
	}
	if res != nil {
		t.Fatalf(GOT, res, WANTED, nil)
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf(GOT, e.Code(), WANTED, twirp.InvalidArgument)
	}
	if e.Msg() != string(emsg) {
		t.Fatalf(GOT, e.Msg(), WANTED, emsg)
	}

}
