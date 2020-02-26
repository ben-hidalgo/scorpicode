package hatserver

import (
	"backend/pkg/util"
	"backend/rpc/hatspb"
)

// Server implements the Hats interface
type Server struct{}

var _ hatspb.Hats = (*Server)(nil)

// NewServer returns an instance of hatserver.Server
func NewServer() *Server {
	return &Server{}
}

//// these will be decoded into multi-lingual, user facing error / warning messages in the UI

// HatInchesTooSmall .
const HatInchesTooSmall = util.ErrMsg("hat.inches.toosmall")

// HatInchesTooBig .
const HatInchesTooBig = util.ErrMsg("hat.inches.toobig")

// HatInchesRequired .
const HatInchesRequired = util.ErrMsg("hat.inches.required")

// HatColorRequired .
const HatColorRequired = util.ErrMsg("hat.color.required")

// HatColorDomain .
const HatColorDomain = util.ErrMsg("hat.color.domain")

// HatSizeRequired .
const HatSizeRequired = util.ErrMsg("hat.size.required")

// HatStyleRequired .
const HatStyleRequired = util.ErrMsg("hat.style.required")

// ToStyle converts a string to a Style type
func ToStyle(s string) hatspb.Style {
	return hatspb.Style(hatspb.Style_value[s])
}
