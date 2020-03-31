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

// HatIDRequired .
const HatIDRequired = util.ErrMsg("hat.id.required")

// HatVersionRequired .
const HatVersionRequired = util.ErrMsg("hat.version.required")

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
