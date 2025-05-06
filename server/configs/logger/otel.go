package logger

import (
	"context"

	"go.uber.org/zap/zapcore"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// otelCore 实现 zapcore.Core 接口
type OtelCore struct {
	zapcore.Core
}

func (c *OtelCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *OtelCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	var ctx context.Context

	for _, f := range fields {
		if f.Key == "ctx" {
			if c, ok := f.Interface.(context.Context); ok {
				ctx = c
				break
			}
		}
	}

	if ctx != nil {
		// 获取当前 span，一次请求中，request和response的span是同一个
		span := trace.SpanFromContext(ctx)

		if span.IsRecording() {
			attrs := []attribute.KeyValue{
				attribute.String("level", ent.Level.String()),
				attribute.String("time", ent.Time.String()),
				attribute.String("message", ent.Message),
				attribute.String("Stack", ent.Stack),
			}

			// 添加 zap fields
			for _, field := range fields {
				if field.Key != "ctx" {
					attrs = append(attrs, attribute.String("field."+field.Key, field.String))
				}
			}

			// 根据日志级别处理
			switch ent.Level {
			case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				span.SetStatus(codes.Error, ent.Message)
				span.AddEvent("log", trace.WithAttributes(attrs...))
			default:
				span.AddEvent("log", trace.WithAttributes(attrs...))
			}
		}
	}
	return c.Core.Write(ent, fields)
}

func (c *OtelCore) With(fields []zapcore.Field) zapcore.Core {
	return &OtelCore{
		Core: c.Core.With(fields),
	}
}
