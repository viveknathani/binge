// package shared provides functions for handling the key-value pairs
// in the context variables being used
package shared

import "context"

type requestID string

func WithRequestID(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, requestID("requestId"), reqId)
}

func ExtractRequestID(ctx context.Context) string {

	result := ""
	data := ctx.Value(requestID("requestId"))
	if data != nil {
		result = data.(string)
	}

	return result
}
