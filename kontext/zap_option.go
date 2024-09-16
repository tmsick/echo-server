package kontext

import (
	"context"

	"go.uber.org/zap"
)

func ZapOption(ctx context.Context) zap.Option {
	id := GetRequestID(ctx)
	if id == "" {
		id = "unknown"
	}
	return zap.Fields(zap.String("request_id", id))
}
