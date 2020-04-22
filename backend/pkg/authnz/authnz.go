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
	"time"

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

// HABERDASHER is a mock role for PoC purposes
const HABERDASHER = Role("HABERDASHER")

// VERIFIED means the user has verified their email address
const VERIFIED = Role("VERIFIED") // probably not a role

// Bearer interface for token abstraction
type Bearer interface {
	// has any role provided
	HasRole(r ...Role) bool
	GetEmail() string
	GetRoles() []Role
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

// HasRole returns true if any of the given roles are present
func (bt *BearerToken) HasRole(roles ...Role) bool {

	for _, br := range bt.GetRoles() {
		for _, r := range roles {
			if br == r {
				return true
			}
		}
	}
	return false
}

// GetEmail .
func (bt *BearerToken) GetEmail() string {
	if bt.CC == nil {
		return ""
	}
	return bt.CC.Email
}

// GetRoles .
func (bt *BearerToken) GetRoles() []Role {

	// no custom claims means anonymous i.e. not authenticated / no roles
	if bt.CC == nil {
		return []Role{}
	}

	var roles []Role

	// TODO: if claims include VERIFIED, append it

	for _, r := range bt.CC.Roles {
		roles = append(roles, Role(r))
	}

	return roles
}

// CustomClaims .
type CustomClaims struct {
	jwt.Claims
	Email string `json:"email"`
	Roles []Role `json:"https://scorpicode.com/roles"`
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
		logrus.Tracef("authnz.traceRequest() len(auth) == 0")
		return
	}

	split := strings.Split(auth, " ")
	if len(split) != 2 {
		logrus.Tracef("authnz.traceRequest() len(split) != 2")
		return
	}

	tokenSplit := strings.Split(split[1], ".")
	if len(tokenSplit) != 3 {
		logrus.Tracef("authnz.traceRequest() len(tokenSplit) != 3")
		return
	}

	decoded, err := base64.RawStdEncoding.DecodeString(tokenSplit[1])
	if err != nil {
		logrus.Tracef("authnz.traceRequest() decoded err=%s     %s", err, tokenSplit[1])
		return
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(decoded, &dat); err != nil {
		logrus.Tracef("authnz.traceRequest() unmarshal err=%s", err)
		return
	}

	// key "email_verified" is the email verified from Auth0
	for k, v := range dat {
		logrus.Tracef("authnz.traceRequest() %s: %s", k, v)
	}

}

// CreateSignedJWT .
func CreateSignedJWT(subject string, roles ...Role) (string, error) {

	//TODO: to consider... we're currently signing this with the private key...
	// I believe that we could use the private key to create a different public key
	// for each environment and sign with the public key.
	// This would be in anticipation of using different public keys for different 3rd party
	// integrations with enterprise business partners.
	// Also, by adding the checksum of the public key to the claims, and storing a map in a database table,
	// we could support key rotation for ourselves, (verifying with the private key)
	// and enable/disable third parties by marking the db column disabled.
	// This could extend to using different public keys in different data centers but
	// they would still be cross-compatible with endpoints at other data centers.

	sig := jose.SigningKey{
		Algorithm: jose.HS256,
		// Key:       privateKeyBytes,
	}

	opts := (&jose.SignerOptions{}).WithType("JWT")

	signer, err := jose.NewSigner(sig, opts)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiry := now.Add(time.Duration(JWTExpirationHours) * time.Hour)

	claims := jwt.Claims{
		Subject:   subject,
		Issuer:    Auth0Issuer,
		Audience:  strings.Split(Auth0Audience, ","),
		NotBefore: jwt.NewNumericDate(now),
		Expiry:    jwt.NewNumericDate(expiry),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	custom := CustomClaims{
		Claims: claims,
		// Roles:  rolez,
	}

	compact, err := jwt.Signed(signer).Claims(custom).CompactSerialize()
	if err != nil {
		return "", err
	}

	return compact, nil
}
