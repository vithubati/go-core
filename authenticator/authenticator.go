package authenticator

import "github.com/go-resty/resty/v2"

// Authenticator describes set of methods implemented by each authenticator
type Authenticator interface {
	AuthenticationType() string
	Authenticate(*resty.Request)
	Validate() error
}
