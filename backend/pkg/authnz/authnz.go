package authnz

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
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

// ValidateRequest .
func ValidateRequest(r *http.Request) (Bearer, error) {

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
	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, err
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

// cache in memory
var pemFile []byte

func loadPublicKey() (interface{}, error) {

	// forward declare the error so as to not shadow the package level pemFile contents
	var err error
	if len(pemFile) == 0 {
		pemFile, err = ioutil.ReadFile(Auth0PemfilePath)
	}
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemFile)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert.PublicKey, nil
}
