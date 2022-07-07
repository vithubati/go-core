package authenticator

import "github.com/go-resty/resty/v2"

type NoAuthAuthenticator struct{}

func NewNoAuthAuthenticator() Authenticator {
	return NoAuthAuthenticator{}
}

func (n NoAuthAuthenticator) AuthenticationType() string {
	return AuthTypeNoAuth
}

func (n NoAuthAuthenticator) Authenticate(request *resty.Request) {
}

func (n NoAuthAuthenticator) Validate() error {
	return nil
}
