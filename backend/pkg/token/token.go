package token

import (
	"backend/pkg/httpwrap"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// Bearer interface for token abstraction
type Bearer interface {
	Clone() Bearer
	// HasRole() bool
}

// used to store the Bearer in Context
type key int

// Key in context
var Key key

//GetBearer returns the repo and panics if not found
func GetBearer(ctx context.Context) Bearer {

	switch v := ctx.Value(Key).(type) {
	case Bearer:
		return v
	default:
		panic("GetRepo() no value found")
	}
}

// Hook middleware injects the DB impl
func Hook(tw Bearer) *twirp.ServerHooks {

	hook := &twirp.ServerHooks{}

	hook.RequestReceived = func(ctx context.Context) (context.Context, error) {

		headers := httpwrap.GetHeaders(ctx)

		logrus.Debugf("headers=%#v", headers)

		return context.WithValue(ctx, Key, tw.Clone()), nil
	}

	return hook
}

// JwtToken impl
type JwtToken struct {
}

// enforces the interface is implemented
var _ Bearer = (*JwtToken)(nil)

// Clone .
func (jt *JwtToken) Clone() Bearer {
	return &JwtToken{}
}
