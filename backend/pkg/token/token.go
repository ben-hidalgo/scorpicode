package token

import (
	"backend/internal/roxie/config"
	"context"
	"net/http"

	// jose "gopkg.in/square/go-jose.v2"
	auth0 "github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// Role is user role
type Role string

// CSR is customer service representative
const CSR = Role("CSR")

// Bearer interface for token abstraction
type Bearer interface {
	// has any role provided
	HasRole(r ...Role) bool
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

// BearerToken impl
type BearerToken struct {
	JWT *jwt.JSONWebToken
}

// enforces the interface is implemented
var _ Bearer = (*BearerToken)(nil)

// HasRole .
func (bt *BearerToken) HasRole(r ...Role) bool {
	return true
}

/*
// ParseJWT parses the provided authorization header
func ParseJWT(auth string) (Bearer, error) {

	if len(auth) == 0 {
		return nil, errors.New("missing authorization header")
	}

	split := strings.Split(auth, " ")

	if len(split) != 2 {
		return nil, errors.New("unexpected length after splitting authorization header")
	}

	obj, err := jwt.ParseSigned(split[1])
	if err != nil {
		return nil, err
	}

	// obj.Claims()

	return &BearerToken{
		JWT: obj,
	}, nil
}
*/

// ValidateRequest .
func ValidateRequest(r *http.Request) (Bearer, error) {
	secret := []byte(config.Auth0ClientSecret)
	secretProvider := auth0.NewKeyProvider(secret)
	audience := []string{config.Auth0Audience}

	configuration := auth0.NewConfiguration(secretProvider, audience, config.Auth0Issuer, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	t, err := validator.ValidateRequest(r)
	if err != nil {
		return nil, err
	}
	return &BearerToken{
		JWT: t,
	}, nil
}
