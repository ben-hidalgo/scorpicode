package hatserver_test

import (
	"backend/internal/hats/config"
	"backend/internal/hats/hatserver"
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"testing"

	"github.com/twitchtv/twirp"
)

const (
	DefaultColor  = "red"
	DefaultName   = "cap"
	DefaultInches = int32(10)

	NOT_NIL = "not nil"
	GOT     = "got '%v' %s '%v'"
	WANTED  = "but wanted"
)

func start() (context.Context, *hatserver.Server, *hatspb.MakeHatRequest) {

	ctx := context.WithValue(context.Background(), repo.RepoKey, inmem.NewRepo())

	hs := hatserver.NewServer()

	req := &hatspb.MakeHatRequest{
		Inches: DefaultInches,
		Name:   DefaultName,
		Color:  DefaultColor,
	}

	return ctx, hs, req
}

func TestSuccess(t *testing.T) {

	ctx, hs, req := start()

	res, err := hs.MakeHat(ctx, req)

	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	if res == nil {
		t.Fatalf(GOT, res, WANTED, NOT_NIL)
	}
	if res.GetHat() == nil {
		t.Fatalf(GOT, res.GetHat(), WANTED, NOT_NIL)
	}
	if res.GetHat().GetInches() != DefaultInches {
		t.Fatalf(GOT, res.GetHat().GetInches(), WANTED, DefaultInches)
	}
	if res.GetHat().GetName() != DefaultName {
		t.Fatalf(GOT, res.GetHat().GetName(), WANTED, DefaultName)
	}
	if res.GetHat().GetColor() != DefaultColor {
		t.Fatalf(GOT, res.GetHat().GetColor(), WANTED, DefaultColor)
	}

}

func TestInchesTooSmall(t *testing.T) {

	ctx, hs, req := start()

	req.Inches = config.MinSizeInches - 1

	res, err := hs.MakeHat(ctx, req)

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
	if e.Msg() != string(hatserver.HatInchesTooSmall) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.HatInchesTooSmall)
	}
}

func TestInchesTooBig(t *testing.T) {

	ctx, hs, req := start()

	req.Inches = config.MaxSizeInches + 1

	res, err := hs.MakeHat(ctx, req)

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
	if e.Msg() != string(hatserver.HatInchesTooBig) {
		t.Fatalf(GOT, e.Msg(), WANTED, hatserver.HatInchesTooSmall)
	}
}

func TestInchesRequired(t *testing.T) {

	ctx, hs, req := start()

	req.Inches = 0

	testRequired(t, ctx, hs, req, hatserver.HatInchesRequired)
}

func TestColorRequired(t *testing.T) {

	ctx, hs, req := start()

	req.Color = ""

	testRequired(t, ctx, hs, req, hatserver.HatColorRequired)
}

func TestNameRequired(t *testing.T) {

	ctx, hs, req := start()

	req.Name = ""

	testRequired(t, ctx, hs, req, hatserver.HatNameRequired)
}

func testRequired(t *testing.T, ctx context.Context, hs *hatserver.Server, req *hatspb.MakeHatRequest, emsg util.ErrMsg) {

	res, err := hs.MakeHat(ctx, req)

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
