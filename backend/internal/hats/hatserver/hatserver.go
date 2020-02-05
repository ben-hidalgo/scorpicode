package hatserver

import (
	"backend/rpc/hatspb"
	"fmt"
)

// Server implements the Hats interface
type Server struct{}

var _ hatspb.Hats = (*Server)(nil)

// twirp meta
const Argument = "argument"

// field names
const Inches = "inches"

// validation messages
const MustBeGTZero = "must be gt zero"

// error messages
var InchesGTZero = emsg(Inches, MustBeGTZero)

func emsg(a, b string) string {
	return fmt.Sprintf("%s %s", a, b)
}
