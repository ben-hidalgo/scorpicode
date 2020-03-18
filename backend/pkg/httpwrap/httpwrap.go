package httpwrap

import (
	"backend/pkg/token"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

// used to store the Headers in the Context
type key int

// Key is the key in context
var Key key

//GetHeaders returns the Headers
func GetHeaders(ctx context.Context) Headers {
	h, ok := ctx.Value(Key).(Headers)
	if !ok {
		return Headers{}
	}
	return h
}

// Headers holds the value of the inbound HTTP headers
type Headers struct {
	Authorization string
	UserAgent     string
	ContentType   string
}

// WithHeaders forwards header from http request to twirp context
func WithHeaders(base http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			return
		}

		// parsing the token here so that a 401 can be returned
		// twirp server hooks don't have access to the response writer
		bearer, err := token.ValidateRequest(r)
		if err != nil {
			logrus.Warnf("WithHeaders() err=%#v", err)
			// any error in the token is a 401
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), token.Key, bearer)
		r = r.WithContext(ctx)

		///////////////////////
		///////////////////////
		headers := &Headers{
			Authorization: r.Header.Get("Authorization"),
			UserAgent:     r.Header.Get("User-Agent"),
			ContentType:   r.Header.Get("Content-Type"),
		}

		ctx = context.WithValue(r.Context(), Key, *headers)
		r = r.WithContext(ctx)

		base.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
