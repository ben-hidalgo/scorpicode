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

// HatVersionMismatch .
const HatVersionMismatch = util.ErrMsg("hat.version.mismatch")

// HatColorRequired .
const HatColorRequired = util.ErrMsg("hat.color.required")

// HatColorDomain .
const HatColorDomain = util.ErrMsg("hat.color.domain")

// HatSizeRequired .
const HatSizeRequired = util.ErrMsg("hat.size.required")

// HatStyleRequired .
const HatStyleRequired = util.ErrMsg("hat.style.required")

// HatQuantityInvalid .
const HatQuantityInvalid = util.ErrMsg("hat.quantity.invalid")

// MakeHatsForbidden .
const MakeHatsForbidden = util.ErrMsg("makehats.forbidden")
