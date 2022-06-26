package authenticator

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBasicAuthenticator(t *testing.T) {
	bAuth := &basicAuthenticator{
		Username: "",
		Password: "password",
	}
	err := bAuth.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf(ErrorMsgProMissing, "Username").Error(), err.Error())

	bAuth = &basicAuthenticator{
		Username: "username",
		Password: "Username",
	}
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf(ErrorMsgProMissing, "Username").Error(), err.Error())
	authenticator, err := NewBasicAuthenticator("username", "password")
	assert.NotNil(t, authenticator)
	assert.Nil(t, err)
}

func TestBasicAuthAuthenticate(t *testing.T) {
	auth, _ := NewBasicAuthenticator("foo", "bar")
	assert.NotNil(t, auth)
	assert.Equal(t, auth.AuthenticationType(), AuthTypeBasic)
	// Create a new Request object
	request := NewRequest("GET", "http://localhost", nil, nil)
	assert.NotNil(t, request)

	// Test the "Authenticate" method to make sure the correct header is added to the Request.
	auth.Authenticate(request)
	assert.ObjectsAreEqualValues(auth, request.UserInfo)
	_, _ = request.Send()
	assert.Equal(t, "Basic Zm9vOmJhcg==", request.Header.Get("Authorization"))
}

// NewRequest will build a new resty request with the given url, body and headers
func NewRequest(method, url string, body interface{}, headers http.Header) *resty.Request {
	r := resty.New().R()
	r.Method = method
	r.URL = url
	if headers != nil {
		r.Header = headers
	}

	r.SetBody(body)
	return r
}
