package hatserver

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"time"
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

// HatColorInvalid .
const HatColorInvalid = util.ErrMsg("hat.color.invalid")

// HatStyleInvalid .
const HatStyleInvalid = util.ErrMsg("hat.style.invalid")

// HatSizeRequired .
const HatSizeRequired = util.ErrMsg("hat.size.required")

// HatStyleRequired .
const HatStyleRequired = util.ErrMsg("hat.style.required")

// HatQuantityInvalid .
const HatQuantityInvalid = util.ErrMsg("hat.quantity.invalid")

// MakeHatsForbidden .
const MakeHatsForbidden = util.ErrMsg("makehats.forbidden")

// HatDocToRep convert Hat document (Mongo) to Hat representation (gRPC)
func HatDocToRep(hat *hatdao.Hat) *hatspb.Hat {
	return &hatspb.Hat{
		Id:        hat.ID.Hex(),
		CreatedAt: hat.CreatedAt.Format(time.RFC3339),
		UpdatedAt: hat.UpdatedAt.Format(time.RFC3339),
		Version:   int32(hat.Version),
		Color:     hat.Color,
		Style:     hat.Style,
		Size:      hat.Size,
		// TODO: hat.Ordinal
		// TODO: hat.OrderID
	}
}

// OrderDocToRep convert Order document (Mongo) to Order representation (gRPC)
func OrderDocToRep(order *orderdao.Order) *hatspb.Order {
	return &hatspb.Order{
		Id:        order.ID.Hex(),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
		Version:   int32(order.Version),
		Color:     order.Color,
		Style:     order.Style,
		Size:      order.Size,
		Quantity:  int32(order.Quantity),
		Notes:     order.Notes,
	}
}

var colors = map[string]interface{}{
	"RED":    struct{}{},
	"BLUE":   struct{}{},
	"GREEN":  struct{}{},
	"YELLOW": struct{}{},
	"PURPLE": struct{}{},
	"BLACK":  struct{}{},
	"GREY":   struct{}{},
	"ORANGE": struct{}{},
}

var styles = map[string]interface{}{
	"BOWLER":   struct{}{},
	"FEDORA":   struct{}{},
	"BASEBALL": struct{}{},
	"NEWSBOY":  struct{}{},
	"COWBOY":   struct{}{},
	"DERBY":    struct{}{},
	"TOP_HAT":  struct{}{},
}
