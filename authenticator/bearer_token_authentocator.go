package authenticator

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// .bearerTokenAuthenticator will take a user-supplied bearer token and adds
// it to requests via an Authorization header of the form:
//
//	Authorization: Bearer <bearer-token>
type bearerTokenAuthenticator struct {
	// auth token scheme type in the HTTP request.For Example:
	//	Authorization: <auth-scheme-value-set-here> <auth-token-value>
	AuthScheme string

	// The bearer token value to be used to authenticate request [required].
	BearerToken string
}

// NewBearerTokenAuthenticator constructs a new bearerTokenAuthenticator instance.
func NewBearerTokenAuthenticator(bearerToken string) (Authenticator, error) {
	auth := &bearerTokenAuthenticator{
		AuthScheme:  "Bearer",
		BearerToken: bearerToken,
	}
	if err := auth.Validate(); err != nil {
		return nil, err
	}
	return auth, nil
}

// AuthenticationType returns authentication type for this authenticator
func (*bearerTokenAuthenticator) AuthenticationType() string {
	return AuthTypeBearerToken
}

// Authenticate adds bearer authentication information to the request.
// The bearer token will be added to the request's headers in the form:
// Authorization: Bearer <bearer-token>
func (a *bearerTokenAuthenticator) Authenticate(request *resty.Request) {
	request.SetAuthScheme(a.AuthScheme)
	request.SetAuthToken(a.BearerToken)
}

// Validate the authenticator's configuration.
// Ensures the bearer token is not empty
func (a *bearerTokenAuthenticator) Validate() error {
	if a.BearerToken == "" {
		return fmt.Errorf(ErrorMsgProMissing, "Bearer Token")
	}
	return nil
}
