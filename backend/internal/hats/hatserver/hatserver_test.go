package hatserver_test

import (
	"backend/internal/hats/hatserver"
	"backend/rpc/hatspb"
	"testing"
)

const (
	NOT_NIL   = "not nil"
	NOT_EMPTY = "not empty"
	GOT       = "got '%v' %s '%v'"
	WANTED    = "but wanted"
)

func TestToStyle(t *testing.T) {

	styletests := []struct {
		in  string
		out hatspb.Style
	}{
		{"BOWLER", hatspb.Style_BOWLER},
		{"FEDORA", hatspb.Style_FEDORA},
		{"BASEBALL", hatspb.Style_BASEBALL},
		{"NEWSBOY", hatspb.Style_NEWSBOY},
		{"COWBOY", hatspb.Style_COWBOY},
		{"DERBY", hatspb.Style_DERBY},
		{"TOP_HAT", hatspb.Style_TOP_HAT},
		{"UNKNOWN_STYLE", hatspb.Style_UNKNOWN_STYLE},
		{"not a style", hatspb.Style_UNKNOWN_STYLE},
	}

	for _, st := range styletests {
		s := hatserver.ToStyle(st.in)
		if s != st.out {
			t.Fatalf(GOT, s, WANTED, st.out)
		}
	}
}
