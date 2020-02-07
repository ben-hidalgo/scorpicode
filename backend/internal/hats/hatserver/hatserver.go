package hatserver

import (
	"backend/pkg/util"
	"backend/rpc/hatspb"
)

// Server implements the Hats interface
type Server struct{}

var _ hatspb.Hats = (*Server)(nil)

// field names
const Inches = "inches"

// these will be decoded into multi-lingual using facing error / warning messages in the UI
const HatTooSmall = util.ErrMsg("hat.inches.toosmall")
const HatTooBig = util.ErrMsg("hat.inches.toobig")
