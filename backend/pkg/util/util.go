package util

import (
	"github.com/twitchtv/twirp"
)

// ErrMsg strong typed error message for constants
type ErrMsg string

// InvalidArgumentError wraps twirp.InvalidArgumentError
func InvalidArgumentError(arg string, emsg ErrMsg) twirp.Error {
	err := twirp.NewError(twirp.InvalidArgument, string(emsg))
	return err
}
