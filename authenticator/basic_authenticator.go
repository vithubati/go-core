package authenticator

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// basicAuthenticator takes a user-supplied username and password, and adds
// them to requests via an Authorization header of the form:
//
//			Authorization: Basic <encoded username and password>
//
type basicAuthenticator struct {
	// Username is the user-supplied basic auth username [required].
	Username string

	//Password is the user-supplied basic auth password [required].
	Password string
}

// NewBasicAuthenticator constructs a new basicAuthenticator instance.
func NewBasicAuthenticator(username string, password string) (Authenticator, error) {
	auth := &basicAuthenticator{
		Username: username,
		Password: password,
	}
	if err := auth.Validate(); err != nil {
		return nil, err

	}
	return auth, nil
}

// AuthenticationType returns the authentication type for this authenticator,
func (basicAuthenticator) AuthenticationType() string {
	return AuthTypeBasic
}

// Authenticate adds basic authentication information to a request.
//
//Basic Authorization will be added to the request's headers in the form:
//
//		Authorization: Basic encoded username and password>
func (a *basicAuthenticator) Authenticate(request *resty.Request) {
	request.SetBasicAuth(a.Username, a.Password)
}

// Validate the authenticator's configuration.
//
// Ensures the username and password are not Nil. Additionally, ensures
// they do not contain invalid characters.
func (a *basicAuthenticator) Validate() error {
	if a.Username == "" {
		return fmt.Errorf(ErrorMsgProMissing, "Username")
	}

	if a.Password == "" {
		return fmt.Errorf(ErrorMsgProMissing, "Password")
	}
	return nil
}
