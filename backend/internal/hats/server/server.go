package server

import "backend/rpc/hatspb"

// Server implements the Hats interface
type Server struct{}

var _ hatspb.Hats = (*Server)(nil)
