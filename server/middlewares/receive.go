package middlewares

import (
	"encoding/json"
	"html"
	"io"
	"strings"

	"gin-server/server/configs/env"
	"gin-server/server/configs/logger"
	"gin-server/server/globals"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

/*收到请求设置trace日志*/
func AcceptRequestHandle(tp *sdktrace.TracerProvider) gin.HandlerFunc {
	if tp != nil {
		otel.SetTracerProvider(tp)
		propagator := propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		)
		otel.SetTextMapPropagator(propagator)
	}

	return func(c *gin.Context) {
		headerMap := make(map[string]string)
		queryMap := make(map[string]string)
		paramMap := make(map[string]string)
		bodyMap := make(map[string]interface{})

		for k := range c.Request.Header {
			if strings.EqualFold(k, "Authorization") || strings.EqualFold(k, "Token") {
				headerMap[k] = "***REDACTED***"
			} else {
				headerMap[k] = c.GetHeader(k)
			}
		}
		for k := range c.Request.URL.Query() {
			queryMap[k] = c.Query(k)
		}
		for _, param := range c.Params {
			paramMap[param.Key] = param.Value
		}

		if strings.HasPrefix(c.ContentType(), "application/json") {
			var bodyBytes []byte

			if c.Request.Body != nil {
				maxBodySize := int64(10 << 20) // 10MB
				_bodyBytes, e := io.ReadAll(io.LimitReader(c.Request.Body, maxBodySize))

				if e != nil {
					panic(globals.NewException(e.Error()))
				}
				bodyBytes = _bodyBytes
				c.Request.Body = io.NopCloser(strings.NewReader(string(_bodyBytes)))
			}
			bodyJsonString := strings.Trim(string(bodyBytes), "\n")

			if er := json.Unmarshal([]byte(bodyJsonString), &bodyMap); er != nil {
				panic(globals.NewException(er.Error()))
			}
		} else if c.ContentType() == "application/x-www-form-urlencoded" || c.ContentType() == "multipart/form-data" {
			if err := c.Request.ParseForm(); err == nil {
				bodyMap = make(map[string]interface{})

				for k, v := range c.Request.PostForm {
					if len(v) == 1 {
						bodyMap[k] = v[0]
					} else {
						bodyMap[k] = v
					}
				}
			}
		}

		type RequestDataType struct {
			Header map[string]string      `json:"header"`
			Query  map[string]string      `json:"query"`
			Body   map[string]interface{} `json:"body"`
			Param  map[string]string      `json:"param"`
		}

		requestData := RequestDataType{
			Header: headerMap,
			Query:  queryMap,
			Body:   bodyMap,
			Param:  paramMap,
		}
		requestDataJson, err := json.MarshalIndent(requestData, "", "    ")

		if err != nil {
			panic(globals.NewException(err.Error()))
		}

		logMsg := "[http-request] " + c.Request.Method + ": " + html.EscapeString(c.Request.URL.RequestURI()) + " " + string(requestDataJson)
		otelEnabled, _ := env.GetEnv("OTEL_ENABLED")

		if otelEnabled == "yes" {
			// 从请求头提取 TraceContext（兼容分布式追踪）
			ctx := otel.GetTextMapPropagator().Extract(
				c.Request.Context(),
				propagation.HeaderCarrier(c.Request.Header),
			)
			// 为每个请求创建根 Span
			ctx, span := otel.Tracer(globals.SERVER_NAME).Start(
				ctx,
				"http-request",
				trace.WithSpanKind(trace.SpanKindServer),
			)
			defer span.End()
			c.Request = c.Request.WithContext(ctx)

			logger.TraceLog.Info(logMsg, zap.Any("ctx", c.Request.Context()))
		} else {
			logger.TraceLog.Info(logMsg)
		}

		c.Set("query", queryMap)
		c.Set("body", bodyMap)
		c.Set("param", paramMap)
		c.Next()
	}
}
