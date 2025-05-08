package middlewares

import (
	"encoding/json"
	"html"
	"io"
	"strings"

	"gin-server/server/configs/logger"
	"gin-server/server/globals"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

/*收到请求设置trace日志*/
func AcceptRequestHandle(c *gin.Context) {
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

	span := trace.SpanFromContext(c.Request.Context())

	if span.IsRecording() {
		span.SetName(c.Request.Method + ": " + c.FullPath())
		span.AddEvent("http-request", trace.WithAttributes(
			attribute.String("log", logMsg),
		))

		spanContext := span.SpanContext()
		logMsg = "[" + spanContext.TraceID().String() + "|" + spanContext.SpanID().String() + "|] " + logMsg
	}

	logger.TraceLog.Info(logMsg)

	c.Set("query", queryMap)
	c.Set("body", bodyMap)
	c.Set("param", paramMap)
	c.Next()
}
