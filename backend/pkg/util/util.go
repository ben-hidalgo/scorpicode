package util

import (
	"github.com/twitchtv/twirp"
)

type ErrMsg string

func InvalidArgumentError(arg string, emsg ErrMsg) twirp.Error {
	err := twirp.NewError(twirp.InvalidArgument, string(emsg))
	return err
}
