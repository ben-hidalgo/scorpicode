package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/rpc/hatspb"
	"context"
	"testing"
)

func startSize() (context.Context, *hatserver.Server, *hatspb.ListSizesRequest) {

	ctx := context.Background()

	hs := hatserver.NewServer()

	req := &hatspb.ListSizesRequest{}

	return ctx, hs, req
}

func TestSizeSuccess(t *testing.T) {

	ctx, hs, req := startSize()

	resp, err := hs.ListSizes(ctx, req)
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}

	for _, s := range resp.GetSizes() {
		if s.GetSlug() == "" {
			t.Fatalf(GOT, s.GetSlug(), WANTED, NOT_EMPTY)
		}
		if s.GetName() == "" {
			t.Fatalf(GOT, s.GetName(), WANTED, NOT_EMPTY)
		}
	}

}
