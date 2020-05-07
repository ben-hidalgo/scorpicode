package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/internal/hats/orderdao"
	"backend/internal/hats/orderdao/orderdaomock"
	"backend/pkg/authnz"
	"backend/pkg/rabbit"
	"backend/pkg/rabbit/rabbitmock"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultColor    = "RED"
	DefaultStyle    = "FEDORA"
	DefaultSize     = "06000"
	DefaultQuantity = 10
	DefaultNotes    = "Lorem ipsum"
	DefaultHexID    = "5e8e20bbe6b38b8cb0870808"
)

var DefaultCreatedAt = time.Date(2020, time.January, 0, 0, 0, 0, 0, time.UTC)
var DefaultUpdatedAt = time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)

func startHat(mock *orderdaomock.Mock, rm *rabbitmock.Mock) (context.Context, *hatserver.Server, *hatspb.MakeHatsRequest) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, orderdao.Key, mock)

	ctx = context.WithValue(ctx, rabbit.Key, rm)

	ctx = context.WithValue(ctx, authnz.Key, &authnz.BearerToken{
		CC: &authnz.CustomClaims{
			Roles: []authnz.Role{
				authnz.HABERDASHER,
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

	mock := &orderdaomock.Mock{
		CreateF: func(ctx context.Context, o *orderdao.Order) error {
			id, err := primitive.ObjectIDFromHex(DefaultHexID)
			if err != nil {
				t.Fatal(err)
			}
			o.SetID(id)
			o.CreatedAt = DefaultCreatedAt
			o.UpdatedAt = DefaultUpdatedAt
			o.Version = 1
			return nil
		},
	}

	rm := &rabbitmock.Mock{
		SendJSONF: func(ex rabbit.Exchange, key rabbit.RKey, msg interface{}) error {
			if ex != rabbit.ServiceMsgtypeTx {
				t.Fatalf(GOT, ex, WANTED, rabbit.ServiceMsgtypeTx)
			}
			if key != rabbit.HatsDotMakeHats {
				t.Fatalf(GOT, key, WANTED, rabbit.HatsDotMakeHats)
			}
			// type assertion
			order, ok := msg.(*orderdao.Order)
			if !ok {
				t.Fatalf(GOT, fmt.Sprintf("%T", msg), WANTED, "*orderdao.Order")
			}
			if order.ID.Hex() != DefaultHexID {
				t.Fatalf(GOT, order.ID.Hex(), WANTED, DefaultHexID)
			}
			if order.CreatedAt != DefaultCreatedAt {
				t.Fatalf(GOT, order.CreatedAt, WANTED, DefaultCreatedAt)
			}
			if order.UpdatedAt != DefaultUpdatedAt {
				t.Fatalf(GOT, order.UpdatedAt, WANTED, DefaultUpdatedAt)
			}
			if order.Version != 1 {
				t.Fatalf(GOT, order.Version, WANTED, 1)
			}
			if order.Size != DefaultSize {
				t.Fatalf(GOT, order.Size, WANTED, DefaultSize)
			}
			if order.Style != DefaultStyle {
				t.Fatalf(GOT, order.Style, WANTED, DefaultStyle)
			}
			if order.Color != DefaultColor {
				t.Fatalf(GOT, order.Color, WANTED, DefaultColor)
			}

			return nil
		},
	}

	ctx, hs, req := startHat(mock, rm)

	res, err := hs.MakeHats(ctx, req)

	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	if res == nil {
		t.Fatalf(GOT, res, WANTED, NOT_NIL)
	}
	if res.GetOrder() == nil {
		t.Fatalf(GOT, res.GetOrder(), WANTED, NOT_NIL)
	}
	if res.GetOrder().GetId() != DefaultHexID {
		t.Fatalf(GOT, res.GetOrder().GetId(), WANTED, DefaultHexID)
	}
	if res.GetOrder().GetCreatedAt() != DefaultCreatedAt.Format(time.RFC3339) {
		t.Fatalf(GOT, res.GetOrder().GetCreatedAt(), WANTED, DefaultCreatedAt.Format(time.RFC3339))
	}
	if res.GetOrder().GetUpdatedAt() != DefaultUpdatedAt.Format(time.RFC3339) {
		t.Fatalf(GOT, res.GetOrder().GetUpdatedAt(), WANTED, DefaultUpdatedAt.Format(time.RFC3339))
	}
	if res.GetOrder().GetVersion() != 1 {
		t.Fatalf(GOT, res.GetOrder().GetVersion(), WANTED, 1)
	}
	if res.GetOrder().GetSize() != DefaultSize {
		t.Fatalf(GOT, res.GetOrder().GetSize(), WANTED, DefaultSize)
	}
	if res.GetOrder().GetStyle() != DefaultStyle {
		t.Fatalf(GOT, res.GetOrder().GetStyle(), WANTED, DefaultStyle)
	}
	if res.GetOrder().GetColor() != DefaultColor {
		t.Fatalf(GOT, res.GetOrder().GetColor(), WANTED, DefaultColor)
	}

}

func TestMissingRole(t *testing.T) {

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

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

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

	req.Size = ""

	testRequired(t, ctx, hs, req, hatserver.HatSizeRequired)
}

func TestColorRequired(t *testing.T) {

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

	req.Color = ""

	testRequired(t, ctx, hs, req, hatserver.HatColorRequired)
}

func TestColorInvalid(t *testing.T) {

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

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
	if e.Msg() != string(hatserver.HatColorInvalid) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.HatColorInvalid)
	}
}

func TestStyleInvalid(t *testing.T) {

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

	req.Style = "not a style"

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
	if e.Msg() != string(hatserver.HatStyleInvalid) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.HatStyleInvalid)
	}
}

func TestNameRequired(t *testing.T) {

	ctx, hs, req := startHat(orderdaomock.New(), rabbitmock.New())

	req.Style = ""

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
