package logger

import (
	"context"

	"github.com/blendle/zapdriver"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	callerSkip1 := zap.AddCallerSkip(1)
	cfg := zapdriver.NewProductionConfig()
	logger = zap.Must(cfg.Build(callerSkip1, zap.AddStacktrace(zapcore.DPanicLevel)))
}

func Sync() {
	logger.Sync()
}
func Info(ctx context.Context, msg string, zapField ...zap.Field) {
	logger.Info(msg, addTraceFromCtx(ctx, zapField)...)
}

func Error(ctx context.Context, msg string, zapField ...zap.Field) {
	logger.Error(msg, addTraceFromCtx(ctx, zapField)...)
}

func Warn(ctx context.Context, msg string, zapField ...zap.Field) {
	logger.Warn(msg, addTraceFromCtx(ctx, zapField)...)
}

func Fatal(ctx context.Context, msg string, zapField ...zap.Field) {
	logger.Fatal(msg, addTraceFromCtx(ctx, zapField)...)
}

func Panic(ctx context.Context, msg string, zapField ...zap.Field) {
	logger.Panic(msg, addTraceFromCtx(ctx, zapField)...)
}

type traceInfo struct {
	traceId  string
	spanId   string
	isSample bool
}

func getTraceFromCtx(ctx context.Context) (isSpanCtxValid bool, t traceInfo) {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		t.traceId = spanContext.TraceID().String()
		t.spanId = spanContext.SpanID().String()
		t.isSample = spanContext.TraceFlags().IsSampled()
	}
	return spanContext.IsValid(), t

}

func addTraceFromCtx(ctx context.Context, fields []zapcore.Field) []zapcore.Field {
	isSpanCtxValid, t := getTraceFromCtx(ctx)
	if isSpanCtxValid {
		fields = append(fields, zapdriver.TraceContext(t.spanId, t.spanId, t.isSample, "ms-ai-marketplace")...)
		fields = append(fields, zap.String("trace_id", t.traceId), zap.Any("span_id", t.spanId))
	}
	return fields
}
