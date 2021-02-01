package sequence

import (
	"bytes"
	"context"
	"time"

	"github.com/segmentio/ksuid"
)

func GenTraceID() string {
	var buf bytes.Buffer

	now := time.Now()
	buf.WriteString(now.Format("20060102"))

	id := ksuid.New()
	buf.WriteString(id.String())
	return buf.String()
}

type ctxKey struct{}

var traceCtxKey ctxKey

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceCtxKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	traceID := ctx.Value(traceCtxKey)
	if val, ok := traceID.(string); ok {
		return val
	}
	return ""
}
