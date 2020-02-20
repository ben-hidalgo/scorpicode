package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"backend/rpc/hatspb"
	"context"
	"testing"

	"github.com/twitchtv/twirp"
)

const (
	XpctdColor = "red"
	expColor   = "red"
	expName    = "cap"
	expInches  = int32(10)
	EXPECTED   = "expected %v %s %v"
	BUT_WAS    = "but was"
	NOT_NIL    = "!nil"
	NOT_EMPTY  = "not empty"
)

func start() (context.Context, *hatserver.Server, *hatspb.MakeHatRequest) {

	ctx := context.WithValue(context.Background(), repo.RepoKey, inmem.NewRepo())

	hs := hatserver.NewServer()

	req := &hatspb.MakeHatRequest{
		Inches: expInches,
		Name:   expName,
		Color:  expColor,
	}

	return ctx, hs, req
}

/*
func TestRandomNameColor(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, repo.RepoKey, inmem.NewRepo())

	inches := int32(10)

	req := &hatspb.MakeHatRequest{
		Inches: inches,
	}

	hs := &hatserver.Server{}

	res, err := hs.MakeHat(ctx, req)

	if err != nil {
		t.Fatalf("err=%s", err)
	}
	if res == nil {
		t.Fatalf("res must not be nil")
	}
	if res.GetHat() == nil {
		t.Fatalf("hat must not be nil")
	}
	if res.GetHat().GetInches() != inches {
		t.Fatalf("inches must match")
	}
	if res.GetHat().GetName() == "" {
		t.Fatalf("name must be randomly assigned when not provided")
	}
	if res.GetHat().GetColor() == "" {
		t.Fatalf("color must be randomly assigned when not provided")
	}
}
*/

func TestSpecificNameColor(t *testing.T) {

	ctx, hs, req := start()

	res, err := hs.MakeHat(ctx, req)

	if err != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, err)
	}
	if res == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if res.GetHat() == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if res.GetHat().GetInches() != expInches {
		t.Fatalf(EXPECTED, expInches, BUT_WAS, res.GetHat().GetInches())
	}
	if res.GetHat().GetName() != expName {
		t.Fatalf(EXPECTED, expName, BUT_WAS, res.GetHat().GetName())
	}
	if res.GetHat().GetColor() != expColor {
		t.Fatalf(EXPECTED, expColor, BUT_WAS, res.GetHat().GetColor())
	}

}

func TestInchesTooSmall(t *testing.T) {

	ctx, hs, req := start()

	req.Inches = int32(4)

	res, err := hs.MakeHat(ctx, req)

	if err == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if res != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, res)
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf(EXPECTED, twirp.InvalidArgument, BUT_WAS, e.Code())
	}
	if e.Msg() != string(hatserver.HatInchesTooSmall) {
		t.Fatalf(EXPECTED, hatserver.HatInchesTooSmall, BUT_WAS, e.Msg())
	}

}

func TestInchesTooBig(t *testing.T) {

	ctx, hs, req := start()

	req.Inches = int32(16)

	res, err := hs.MakeHat(ctx, req)

	if err == nil {
		t.Fatalf(EXPECTED, NOT_NIL, BUT_WAS, nil)
	}
	if res != nil {
		t.Fatalf(EXPECTED, nil, BUT_WAS, res)
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf(EXPECTED, twirp.InvalidArgument, BUT_WAS, e.Code())
	}
	if e.Msg() != string(hatserver.HatInchesTooBig) {
		t.Fatalf(EXPECTED, hatserver.HatInchesTooSmall, BUT_WAS, e.Msg())
	}
}
