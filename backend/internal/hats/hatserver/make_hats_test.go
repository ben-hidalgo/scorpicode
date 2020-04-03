package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"testing"

	"github.com/twitchtv/twirp"
)

const (
	DefaultColor = "RED"
	DefaultStyle = hatspb.Style_FEDORA
	DefaultSize  = "06000"
)

func startHat() (context.Context, *hatserver.Server, *hatspb.MakeHatsRequest) {

	//ctx := context.WithValue(context.Background(), repo.RepoKey, inmem.NewRepo())
	ctx := context.Background()

	hs := hatserver.NewServer()

	req := &hatspb.MakeHatsRequest{
		Size:  DefaultSize,
		Style: DefaultStyle,
		Color: DefaultColor,
	}

	return ctx, hs, req
}

func TestHatSuccess(t *testing.T) {

	// ctx, hs, req := startHat()

	// res, err := hs.MakeHats(ctx, req)

	// if err != nil {
	// 	t.Fatalf(GOT, err, WANTED, nil)
	// }
	// if res == nil {
	// 	t.Fatalf(GOT, res, WANTED, NOT_NIL)
	// }
	// if res.GetHat() == nil {
	// 	t.Fatalf(GOT, res.GetHat(), WANTED, NOT_NIL)
	// }
	// if res.GetHat().GetSize() != DefaultSize {
	// 	t.Fatalf(GOT, res.GetHat().GetSize(), WANTED, DefaultSize)
	// }
	// if res.GetHat().GetStyle() != DefaultStyle {
	// 	t.Fatalf(GOT, res.GetHat().GetStyle(), WANTED, DefaultStyle)
	// }
	// if res.GetHat().GetColor() != DefaultColor {
	// 	t.Fatalf(GOT, res.GetHat().GetColor(), WANTED, DefaultColor)
	// }

}

func TestSizeRequired(t *testing.T) {

	ctx, hs, req := startHat()

	req.Size = ""

	testRequired(t, ctx, hs, req, hatserver.HatSizeRequired)
}

func TestColorRequired(t *testing.T) {

	ctx, hs, req := startHat()

	req.Color = ""

	testRequired(t, ctx, hs, req, hatserver.HatColorRequired)
}

func TestColorDomain(t *testing.T) {

	ctx, hs, req := startHat()

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

	ctx, hs, req := startHat()

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
