package authenticator

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// xTokenAuthenticator will take a user supplied any token header key and token then it
//  adds it to request via Header of the form
//		<HeaderKey>: «Token>
//
type xTokenAuthenticator struct {
	// HeaderKey is the key in the http request Header (required). For example:
	//    «HeaderKey>: <Token Value>
	//
	HeaderKey string

	// The token value to be used to authenticate request (required).
	Token string
}

// NewXTokenAuthenticator constructs a new bearerTokenAuthenticator instance.
func NewXTokenAuthenticator(headerkey, token string) (Authenticator, error) {
	obj := &xTokenAuthenticator{
		HeaderKey: headerkey,
		Token:     token,
	}

	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj, nil
}

// AuthenticationType returns the authentication type for this authenticator.
func (*xTokenAuthenticator) AuthenticationType() string {
	return AuthTypeXToken
}

// Authenticate adds token authentication information to the request
// The token will be added to the request's headers in the forme
//
//		«HeaderKey>: <token value>
//
func (a *xTokenAuthenticator) Authenticate(request *resty.Request) {
	request.SetHeader(a.HeaderKey, a.Token)

}

// Validate the authenticator's configuration.
// Ensures the HeaderKey and token are not empty
func (a *xTokenAuthenticator) Validate() error {
	if len(a.HeaderKey) == 0 {
		return fmt.Errorf(ErrorMsgProMissing, "HeaderKey")
	}
	if len(a.Token) == 0 {
		return fmt.Errorf(ErrorMsgProMissing, "Token")
	}
	return nil
}
