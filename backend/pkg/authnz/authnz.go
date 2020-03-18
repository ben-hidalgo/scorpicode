package authnz

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
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

var pemFile []byte

// ValidateRequest .
func ValidateRequest(r *http.Request) (Bearer, error) {

	// the auth0 internals do not validate inputs
	// if clientSecret == "" {
	// 	return nil, errors.New("validateRequest() clientSecret is required")
	// }
	// if audience == "" {
	// 	return nil, errors.New("validateRequest() audience is required")
	// }
	if Auth0Issuer == "" {
		return nil, errors.New("validateRequest() Auth0Issuer is required")
	}
	if Auth0ClientID == "" {
		return nil, errors.New("validateRequest() Auth0ClientID is required")
	}
	if Auth0PemfilePath == "" {
		return nil, errors.New("validateRequest() Auth0PemfilePath is required")
	}

	var err error

	/////////////////////////////
	if len(pemFile) == 0 {
		pemFile, err = ioutil.ReadFile(Auth0PemfilePath)
		if err != nil {
			return nil, err
		}
	}
	publicKey, err := loadPublicKey(pemFile)
	if err != nil {
		panic(err)
	}

	/////////////////////////////////
	secretProvider := auth0.NewKeyProvider(publicKey)

	// audience is the client ID
	configuration := auth0.NewConfiguration(secretProvider, []string{Auth0ClientID}, Auth0Issuer, jose.RS256)

	validator := auth0.NewValidator(configuration, nil)

	t, err := validator.ValidateRequest(r)
	if err != nil {
		return nil, err
	}

	return &BearerToken{
		JWT: t,
	}, nil
}

func loadPublicKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	// Try to load SubjectPublicKeyInfo
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
		return pub, nil
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	return nil, fmt.Errorf("square/go-jose: parse error, got '%s' and '%s'", err0, err1)
}
