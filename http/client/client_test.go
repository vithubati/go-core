package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vithubati/go-core/authenticator"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	client := New(WithDebug())
	assert.NotNil(t, client)
	auth, err := authenticator.NewBearerTokenAuthenticator("Test-TOKEN")
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	client = New(WithAuthenticator(auth), WithDebug())
	assert.Nil(t, err)
	assert.NotNil(t, auth)
}

func TestRequest(t *testing.T) {
	t.Parallel()
	hUrl := "https://gobyexample.com/"
	auth, _ := authenticator.NewBearerTokenAuthenticator("Test-TOKEN")
	assert.NotNil(t, auth)
	client := New(WithAuthenticator(auth), WithDebug())
	assert.NotNil(t, client)
	req := client.Request(hUrl, http.MethodGet, nil, nil)
	assert.NotNil(t, req)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, hUrl, req.URL)
}

func TestRequestWithCtx(t *testing.T) {
	t.Parallel()
	hUrl := "https://gobyexample.com/"
	client := New(WithDebug())
	ctx := context.Background()
	req := client.RequestWithCtx(ctx, hUrl, http.MethodGet, nil, nil)
	assert.NotNil(t, req)
}

func TestRequestWithID(t *testing.T) {
	t.Parallel()
	hUrl := "https://gobyexample.com/"
	reqId := "cca61568-583-1lec-9d64-0242ac120002"
	client := New(WithDebug())
	req := client.RequestWithID(reqId, hUrl, http.MethodGet, nil, nil)
	assert.NotNil(t, req)
	assert.Equal(t, reqId, req.Header.Get("x-request-id"))
}

func TestExecute(t *testing.T) {
	t.Parallel()
	tUrl := "http://echo.jsontest.com/name/vithu/team/avengers/job/developer"
	reqId := "cca61568-583-1lec-9d64-0242ac120002"
	auth, _ := authenticator.NewBearerTokenAuthenticator("Test-TOKEN")
	client := New(WithAuthenticator(auth))
	params := url.Values{}
	params.Set("limit", "5")
	params.Set("size", "10")
	req := client.RequestWithID(reqId, tUrl, http.MethodGet, nil, params)
	assert.NotNil(t, req)
	resp, err := client.Execute(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "http://echo.jsontest.com/name/vithu/team/avengers/job/developer?limit=5&size=10", resp.Request.URL)
	fmt.Printf("Trace Info: %+v\n", resp.Request.TraceInfo())
}

func TestWithAuthenticator(t *testing.T) {
	auth, _ := authenticator.NewXTokenAuthenticator("TEST-HEADER", "TEST-HEADER-VALUE")
	c := New()
	c.SetOptions(WithAuthenticator(auth))
	r := c.Request("bla", http.MethodGet, nil, nil)
	assert.True(t, strings.Contains(r.Header.Get("TEST-HEADER"), "TEST-HEADER-VALUE"))
}

func TestSetOptions(t *testing.T) {
	auth, _ := authenticator.NewXTokenAuthenticator("TEST-HEADER", "TEST-HEADER-VALUE")
	c := New()
	c.SetOptions(WithDebug(), WithAuthenticator(auth))
	r := c.Request("bla", http.MethodGet, nil, nil)
	assert.True(t, strings.Contains(r.Header.Get("TEST-HEADER"), "TEST-HEADER-VALUE"))
	assert.True(t, c.Debug)
}
