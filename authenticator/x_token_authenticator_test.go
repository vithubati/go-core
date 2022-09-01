package authenticator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewXTokenAuthenticator(t *testing.T) {
	bAuth := &xTokenAuthenticator{
		HeaderKey: "",
		Token:     "f00-token",
	}
	err := bAuth.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf(ErrorMsgProMissing, "HeaderKey").Error(), err.Error())
	bAuth = &xTokenAuthenticator{
		HeaderKey: "HeaderKey",
		Token:     "",
	}
	err = bAuth.Validate()
	assert.Equal(t, fmt.Errorf(ErrorMsgProMissing, "Token").Error(), err.Error())
	authenticator, err := NewXTokenAuthenticator("my-bearer-key", "my-bearer-token")
	assert.NotNil(t, authenticator)
	assert.Nil(t, err)
}

func TestXTokenAuthenticationType(t *testing.T) {
	authenticator, err := NewXTokenAuthenticator("X-AUTH-TOKEN", "my-bearer-token")
	assert.Nil(t, err)
	assert.NotNil(t, authenticator)
	assert.Equal(t, authenticator.AuthenticationType(), AuthTypeXToken)
}

func TestXTokenAuthenticator(t *testing.T) {
	authenticator, _ := NewXTokenAuthenticator("X-AUTH-TOKEN", "my-x-token")
	assert.NotNil(t, authenticator)

	request := NewRequest("GET", "http://localhost", nil, nil)
	assert.NotNil(t, request)

	// Test the "Authenticate" method to make sure the correct header is added to the Request.
	authenticator.Authenticate(request)
	assert.ObjectsAreEqualValues(authenticator, request.UserInfo)
	_, _ = request.Send()
	assert.Equal(t, "my-x-token", request.Header.Get("X-AUTH-TOKEN"))
}
