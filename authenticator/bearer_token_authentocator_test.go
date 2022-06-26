package authenticator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBearerToken(t *testing.T) {
	t.Parallel()
	bAuth := &bearerTokenAuthenticator{
		BearerToken: "",
	}
	err := bAuth.Validate()
	assert.NotNil(t, err)
	assert.Equal(t,
		fmt.Errorf(ErrorMsgProMissing, "Bearer Token").Error(), err.Error())
	authenticator, err := NewBearerTokenAuthenticator("my-bearer-token")
	assert.NotNil(t, authenticator)
	assert.Nil(t, err)
}

func TestBearerTokenAuthenticate(t *testing.T) {
	t.Parallel()
	authenticator, err := NewBearerTokenAuthenticator("my-bearer-token")
	assert.Nil(t, err)
	assert.NotNil(t, authenticator)
	assert.Equal(t, authenticator.AuthenticationType(), AuthTypeBearerToken)

	// Create a new Request object.
	request := NewRequest("GET", "http://localhost", nil, nil)
	assert.NotNil(t, request)

	// Test the "Authenticate" method to make sure the correct header is added to the Request.
	authenticator.Authenticate(request)
	assert.ObjectsAreEqualValues(authenticator, request.UserInfo)
	_, _ = request.Send()
	assert.Equal(t, "Bearer my-bearer-token", request.Header.Get("Authorization"))
}
