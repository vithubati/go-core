package authenticator

import "github.com/go-resty/resty/v2"

// Authenticator describes set of methods implemented by each authenticator
type Authenticator interface {
	// AuthenticationType returns authentication type of this authenticator
	AuthenticationType() string

	// Authenticate - Authenticates the request
	Authenticate(*resty.Request)

	// Validate the authenticator's configuration.
	Validate() error
}
