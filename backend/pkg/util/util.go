package util

import (
	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// ErrMsg strongly typed error message string for constants
type ErrMsg string

// InvalidArgumentError wraps twirp.InvalidArgumentError
func InvalidArgumentError(emsg ErrMsg) twirp.Error {
	logrus.Debugf("InvalidArgumentError emsg=%s", emsg)
	return twirp.NewError(twirp.InvalidArgument, string(emsg))
}

// InternalErrorWith wraps twirp.InternalErrorWith
func InternalErrorWith(e error) twirp.Error {
	logrus.Errorf("InternalErrorWith e=%v", e)
	return twirp.InternalErrorWith(e)
}

// NotFoundError wraps twirp.NotFoundError
func NotFoundError(msg string) twirp.Error {
	logrus.Warnf("NotFoundError msg=%s", msg)
	return twirp.NotFoundError(msg)
}

// PermissionDeniedError wraps twirp.PermissionDeniedError
func PermissionDeniedError(emsg ErrMsg) twirp.Error {
	logrus.Warnf("PermissionDeniedError emsg=%s", emsg)
	return twirp.NewError(twirp.PermissionDenied, string(emsg))
}
