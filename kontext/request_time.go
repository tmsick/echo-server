package kontext

import (
	"context"
	"time"
)

type requestTimeKeyType string

const requestTimeKey requestTimeKeyType = "request_time"

func GetRequestTime(ctx context.Context) time.Time {
	v, ok := ctx.Value(requestTimeKey).(time.Time)
	if !ok {
		return time.Time{}
	}
	return v
}

func SetRequestTime(ctx context.Context, requestTime time.Time) context.Context {
	return context.WithValue(ctx, requestTimeKey, requestTime)
}
