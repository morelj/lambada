package lambada

import (
	"context"
	"net/http"
)

type contextKeyType struct{}

var contextKey = contextKeyType{}

// GetRequest returns the original API Gateway request which issued the http.Request.
// The returned Request value contains both API Gateway V1 and V2 data, but the used fields depend on the actual
// API Gateway version used.
//
// When no API Gateway request is attached to the http.Request, this function returns nil.
func GetRequest(r *http.Request) *Request {
	return GetRequestFromContext(r.Context())
}

// GetRequestFromContext returns the original API Gateway request stored in ctx.
// The returned Request value contains both API Gateway V1 and V2 data, but the used fields depend on the actual
// API Gateway version used.
//
// When no API Gateway request is attached to the http.Request, this function returns nil.
func GetRequestFromContext(ctx context.Context) *Request {
	if res, ok := ctx.Value(contextKey).(*Request); ok {
		return res
	}
	return nil
}

// WithRequest returns a new context.Context with the given Request attached.
// The returned context should then be attached to an http.Request.
//
// There's usually no need to call this function directly, as all the work is done by lambada itself. However, it may
// be useful for testing purposes.
func WithRequest(ctx context.Context, req *Request) context.Context {
	return context.WithValue(ctx, contextKey, req)
}
