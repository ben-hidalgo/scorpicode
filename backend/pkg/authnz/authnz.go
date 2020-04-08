package authnz

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// jose "gopkg.in/square/go-jose.v2"
	auth0 "github.com/auth0-community/go-auth0"
	"github.com/sirupsen/logrus"
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
	GetEmail() string
	GetRoles() []string
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
	CC  *CustomClaims
}

// enforces the interface is implemented
var _ Bearer = (*BearerToken)(nil)

// HasRole .
func (bt *BearerToken) HasRole(r ...Role) bool {
	return true
}

// GetEmail .
func (bt *BearerToken) GetEmail() string {
	if bt.CC == nil {
		return ""
	}
	return bt.CC.Email
}

// GetRoles .
func (bt *BearerToken) GetRoles() []string {
	if bt.CC == nil {
		return []string{}
	}
	return bt.CC.Roles
}

// CustomClaims .
type CustomClaims struct {
	jwt.Claims
	Email string   `json:"email"`
	Roles []string `json:"https://scorpicode.com/roles"`
}

// DEBU[0622] debug() https://scorpicode.com/roles: [CSR]

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

	traceRequest(r)

	// read the pem file and parse the certificate
	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, err
	}

	// the key provider holds the public key value
	secretProvider := auth0.NewKeyProvider(publicKey)

	// audience is the client ID
	configuration := auth0.NewConfiguration(secretProvider, []string{Auth0ClientID}, Auth0Issuer, jose.RS256)

	// nil defaults the extractor to: from request authorization header
	validator := auth0.NewValidator(configuration, nil)

	jt, err := validator.ValidateRequest(r)
	if err != nil {
		return nil, err
	}

	cc := &CustomClaims{}

	err = jt.Claims(publicKey, cc)
	if err != nil {
		return nil, err
	}

	return &BearerToken{
		JWT: jt,
		CC:  cc,
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

	p, rest := pem.Decode(pemFile)
	if len(rest) != 0 {
		return nil, fmt.Errorf("loadPublicKey() unexpected len(rest)=%d", len(rest))
	}

	cert, err := x509.ParseCertificate(p.Bytes)

	if err != nil {
		return nil, err
	}

	return cert.PublicKey, nil
}

func traceRequest(r *http.Request) {

	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		logrus.Tracef("debug() len(auth) == 0")
		return
	}

	split := strings.Split(auth, " ")
	if len(split) != 2 {
		logrus.Tracef("debug() len(split) != 2")
		return
	}

	tokenSplit := strings.Split(split[1], ".")
	if len(tokenSplit) != 3 {
		logrus.Tracef("debug() len(tokenSplit) != 3")
		return
	}

	decoded, err := base64.RawStdEncoding.DecodeString(tokenSplit[1])
	if err != nil {
		logrus.Tracef("debug() decoded err=%s     %s", err, tokenSplit[1])
		return
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(decoded, &dat); err != nil {
		logrus.Tracef("debug() unmarshal err=%s", err)
		return
	}

	for k, v := range dat {
		logrus.Tracef("debug() %s: %s", k, v)
	}

}
