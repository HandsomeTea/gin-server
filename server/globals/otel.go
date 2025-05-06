package globals

import (
	"gin-server/server/configs/env"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitOtelTracer() (*trace.TracerProvider, error) {
	otelEnabled, _ := env.GetEnv("OTEL_ENABLED")

	if otelEnabled != "yes" {
		return nil, nil
	}
	// 创建控制台 Exporter（输出到 stderr）
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(os.Stderr),
		stdouttrace.WithPrettyPrint(), // 美化输出
	)
	if err != nil {
		panic(err)
	}

	// 创建 TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()), // 确保采样
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
