package util

import (
	"github.com/twitchtv/twirp"
)

// twirp meta
const Argument = "argument"

type ErrMsg string

func InvalidArgumentError(arg string, emsg ErrMsg) twirp.Error {
	err := twirp.NewError(twirp.InvalidArgument, string(emsg))
	err = err.WithMeta(Argument, arg)
	return err
}
