package hatserver

import (
	"backend/pkg/util"
	"backend/rpc/hatspb"
)

// Server implements the Hats interface
type Server struct{}

var _ hatspb.Hats = (*Server)(nil)

//// field names

// Inches .
const Inches = "inches"

//// these will be decoded into multi-lingual, user facing error / warning messages in the UI

// HatTooSmall .
const HatTooSmall = util.ErrMsg("hat.inches.toosmall")

// HatTooBig .
const HatTooBig = util.ErrMsg("hat.inches.toobig")
