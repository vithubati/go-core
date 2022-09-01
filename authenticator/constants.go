package authenticator

const (
	//  supported authentication types
	AuthTypeBasic       = "Basic"
	AuthTypeBearerToken = "BearerToken"
	AuthTypeNoAuth      = "NoAuth"
	AuthTypeXToken      = "xToken"

	// common errors
	ErrorMsgProMissing    = "the %s required but was not specified"
	ErrorMsgAuthenticator = "authentication information was not properly configured"
)
