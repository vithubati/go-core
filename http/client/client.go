package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/vithubati/go-core/authenticator"
	"net/url"
)

type Option func(*client)

// Client implements the common http functionality
// to manage requests and responses, authenticate outbound requests, etc.
type client struct {
	// authenticator holds the authenticator implementation to be used by the
	// service instance to authenticate outbound requests, typically by adding the
	// HTTP "Authorization" header.
	authenticator authenticator.Authenticator

	// The HTTP Client used to send requests and receive responses
	*resty.Client
}

// New constructs a new type of Service. it does the validation on the required options
// and returns the New Service
// Parameters:
//		authenticator: holds the authenticator implementation to be used by the
//                     service instance to authenticate outbound requests, typically by adding the
//						HTTP  "Authorization" header.
//
func New(opts ...Option) *client {
	c := &client{
		Client: resty.New(),
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.authenticator == nil {
		c.authenticator = authenticator.NewNoAuthAuthenticator()
	}
	return c
}

// Request Creates a new *resty.Request value from the current Service *resty.Request
// Parameters:
// 			url: full url of the request
// 			method: request method
// 			body: the interface ()  type, holds the request body information - pass nil if its GET request
// Returns: a new *resty.Request value with provided authentication enabled
func (c *client) Request(url, method string, body interface{}, params url.Values) *resty.Request {
	return c.request(url, method, body, params)
}

// RequestWithCtx Creates new *resty.Request value from the current Service *resty.Request
// Parameters:
// 			ctx: full url of the request
//	 		urls full url of the request
// 			method: request method
// 			body: the interface() type, holds the request body information - pass nil if ita GET request
//			params: url.Values set rhe params as url.Value
// Returns: A new *resty.Request value with provided authentication enabled
func (c *client) RequestWithCtx(ctx context.Context, url, method string, body interface{}, params url.Values) *resty.Request {
	return c.request(url, method, body, params).SetContext(ctx)
}

// RequestWithID - sets the provided requestId to the client request header and returns the current Request
func (c *client) RequestWithID(requestId, url, method string, body interface{}, params url.Values) *resty.Request {
	r := c.request(url, method, body, params)
	r.SetHeader("x-request-id", requestId)
	return r
}

//Execute - sends the passed requests and returns the response
func (c *client) Execute(request *resty.Request) (*resty.Response, error) {
	return request.Send()
}

// WithLogger method sets given writer for logging Resty request and response detail
// Compliant to interface  'resty. Logger
func WithLogger(l resty.Logger) Option {
	return func(client *client) {
		client.SetLogger(l)
	}
}

// WithDebug enables debug mode
func WithDebug() Option {
	return func(client *client) {
		client.SetDebug(true)
		client.EnableTrace()
	}
}

// WithAuthenticator set the given authenticator on the client
func WithAuthenticator(auth authenticator.Authenticator) Option {
	return func(client *client) {
		client.authenticator = auth
	}
}

func (c *client) request(url, method string, body interface{}, params url.Values) *resty.Request {
	r := c.Client.R()
	r.Method = method
	r.URL = url
	r.QueryParam = params
	r.Body = body
	c.authenticator.Authenticate(r)
	return r
}
