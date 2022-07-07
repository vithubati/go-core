package authenticator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoAuthAuthenticator(t *testing.T) {
	t.Parallel()
	auth := NewNoAuthAuthenticator()
	assert.NotNil(t, auth)
	assert.Equal(t, AuthTypeNoAuth, auth.AuthenticationType())

	err := auth.Validate()
	assert.Nil(t, err)
}
