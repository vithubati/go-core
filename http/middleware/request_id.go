package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Ported from Goji's middleware, source:
// https://github.com/go-chi/chi/tree/master/middleware/request_id.go
// Key to use when setting the request ID.
type ctxKeyRequestID int

// RequestIDKey is the key that holds the unique request ID in a request context.
const RequestIDKey ctxKeyRequestID = 0

// RequestID is a middleware that injects a request ID into the context of each
// It is using https://github.com/google/uuid to generate UUID v4
// in case of error its generating using the time.Now
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			uuid, err := uuid.NewRandom()
			if err != nil {
				requestID = fmt.Sprint(time.Now().UnixNano())
			} else {
				requestID = uuid.String()
			}
		}
		ctx = context.WithValue(ctx, RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
