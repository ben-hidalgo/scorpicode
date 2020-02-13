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

func TestSpecificNameColor(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, repo.RepoKey, inmem.NewRepo())

	inches := int32(10)
	name := "cap"
	color := "blue"

	req := &hatspb.MakeHatRequest{
		Inches: inches,
		Name:   name,
		Color:  color,
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
	if res.GetHat().GetName() != name {
		t.Fatalf("name must match")
	}
	if res.GetHat().GetColor() != color {
		t.Fatalf("color must match")
	}

}

func TestInchesTooSmall(t *testing.T) {

	ctx := context.Background()

	inches := int32(4)

	req := &hatspb.MakeHatRequest{
		Inches: inches,
	}

	hs := &hatserver.Server{}

	res, err := hs.MakeHat(ctx, req)

	if err == nil {
		t.Fatalf("an error should be returned")
	}
	if res != nil {
		t.Fatalf("the response should be nil")
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf("the error code should be %s", twirp.InvalidArgument)
	}
	if e.Msg() != string(hatserver.HatTooSmall) {
		t.Fatalf("the msg should be '%s' but was '%s'", hatserver.HatTooSmall, e.Msg())
	}

}

func TestInchesTooBig(t *testing.T) {

	ctx := context.Background()

	inches := int32(16)

	req := &hatspb.MakeHatRequest{
		Inches: inches,
	}

	hs := &hatserver.Server{}

	res, err := hs.MakeHat(ctx, req)

	if err == nil {
		t.Fatalf("an error should be returned")
	}
	if res != nil {
		t.Fatalf("the response should be nil")
	}

	e := err.(twirp.Error)
	if e.Code() != twirp.InvalidArgument {
		t.Fatalf("the error code should be %s", twirp.InvalidArgument)
	}
	if e.Msg() != string(hatserver.HatTooBig) {
		t.Fatalf("the msg should be '%s' but was '%s'", hatserver.HatTooBig, e.Msg())
	}
}
