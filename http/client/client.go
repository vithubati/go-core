package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/vithubati/go-core/authenticator"
	"github.com/vithubati/go-core/http/middleware"
	"net/http"
	"net/url"
)

type Option func(*Client)

// Client implements the common http functionality
// to manage requests and responses, authenticate outbound requests, etc.
type Client struct {
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
//
//			authenticator: holds the authenticator implementation to be used by the
//	                    service instance to authenticate outbound requests, typically by adding the
//							HTTP  "Authorization" header.
func New(opts ...Option) *Client {
	c := &Client{
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
//
//	url: full url of the request
//	method: request method
//	body: the interface ()  type, holds the request body information - pass nil if its GET request
//
// Returns: a new *resty.Request value with provided authentication enabled
func (c *Client) Request(url, method string, body interface{}, params url.Values) *resty.Request {
	return c.request(url, method, body, params)
}

// RequestWithCtx Creates new *resty.Request value from the current Service *resty.Request
// Parameters:
//
//			ctx: full url of the request
//	 		urls full url of the request
//			method: request method
//			body: the interface() type, holds the request body information - pass nil if ita GET request
//			params: url.Values set rhe params as url.Value
//
// Returns: A new *resty.Request value with provided authentication enabled
func (c *Client) RequestWithCtx(ctx context.Context, url, method string, body interface{}, params url.Values) *resty.Request {
	r := c.request(url, method, body, params).SetContext(ctx)
	reqID := middleware.GetReqID(ctx)
	if reqID != "" {
		r.SetHeader("x-request-id", reqID)
	}
	return r
}

// RequestWithID - sets the provided requestId to the client request header and returns the current Request
func (c *Client) RequestWithID(requestId, url, method string, body interface{}, params url.Values) *resty.Request {
	r := c.request(url, method, body, params)
	r.SetHeader("x-request-id", requestId)
	return r
}

// Execute - sends the passed requests and returns the response
func (c *Client) Execute(request *resty.Request) (*resty.Response, error) {
	return request.Send()
}

func (c *Client) request(url, method string, body interface{}, params url.Values) *resty.Request {
	r := c.Client.R()
	r.Method = method
	r.URL = url
	if params != nil {
		r.QueryParam = params
	}
	r.Body = body
	c.authenticator.Authenticate(r)
	return r
}

// SetOptions - sets the given client options
func (c *Client) SetOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithLogger method sets given writer for logging Resty request and response detail
// Compliant to interface  'resty. Logger
func WithLogger(l resty.Logger) Option {
	return func(client *Client) {
		client.SetLogger(l)
	}
}

// WithDebug enables debug mode
func WithDebug() Option {
	return func(client *Client) {
		client.SetDebug(true)
		client.EnableTrace()
	}
}

// WithAuthenticator set the given authenticator on the client
func WithAuthenticator(auth authenticator.Authenticator) Option {
	return func(client *Client) {
		client.authenticator = auth
	}
}

// WithTransport set the given http.RoundTripper on the client
func WithTransport(transport http.RoundTripper) Option {
	return func(client *Client) {
		client.SetTransport(transport)
	}
}
