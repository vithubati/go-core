package middleware

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetTestHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		panic("Test entered handler")
	}
}

func TestRequestID(t *testing.T) {
	tests := map[string]struct {
		requestIDHeader  string
		request          func() *http.Request
		expectedResponse string
		random           bool
	}{
		"Retrieves Request Id from default header": {
			"X-Request-Id",
			func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("X-Request-Id", "req-123456")
				return req
			},
			"RequestID: req-123456",
			false,
		},
		"Retrieves Request Id without setting in the header": {
			"X-Request-Id",
			func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				return req
			},
			"",
			true,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		f := func(next http.Handler) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				requestID := GetReqID(r.Context())
				response := fmt.Sprintf("RequestID: %s", requestID)
				w.Write([]byte(response))
			}
		}
		h := RequestID(f(GetTestHandler()))
		h.ServeHTTP(w, test.request())

		if !test.random {
			if w.Body.String() != test.expectedResponse {
				t.Fatalf("RequestID was not the expected value. Expected: %s. Actual: %s", test.expectedResponse, w.Body.String())
			}
		} else {
			if w.Body.String() == "" {
				t.Fatalf("RequestID was not the expected value. Expected: %s. Actual: %s", test.expectedResponse, w.Body.String())
			}
		}
	}
}

func TestGetReqID(t *testing.T) {
	expected := "req-112233"
	ctx := context.Background()
	reqID := GetReqID(ctx)
	if reqID != "" {
		t.Fatalf("RequestID should be empty but it is %s", reqID)
	}
	ctx2 := context.WithValue(ctx, RequestIDKey, expected)
	reqID = GetReqID(ctx2)
	assert.Equal(t, expected, reqID)
}
