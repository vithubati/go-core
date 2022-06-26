package authenticator

const (
	//  supported authentication types
	AuthTypeBasic       = "Basic"
	AuthTypeBearerToken = "bearerToken"

	// common errors
	ErrorMsgProMissing    = "the %s required but was not specified"
	ErrorMsgAuthenticator = "authentication information was not properly configured"
)
