package kontext

import (
	"context"
)

type requestIDKeyType string

const requestIDKey requestIDKeyType = "request_id"

func GetRequestID(ctx context.Context) string {
	v, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}
	return v
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
