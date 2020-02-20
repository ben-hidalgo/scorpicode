package util

import (
	"github.com/twitchtv/twirp"
)

// ErrMsg strongly typed error message string for constants
type ErrMsg string

// InvalidArgumentError wraps twirp.InvalidArgumentError
func InvalidArgumentError(emsg ErrMsg) twirp.Error {
	err := twirp.NewError(twirp.InvalidArgument, string(emsg))
	return err
}
