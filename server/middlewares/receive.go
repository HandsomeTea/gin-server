package middlewares

import (
	"encoding/json"
	"io"
	"strings"

	logger "gin-server/server/configs/logger"

	"github.com/gin-gonic/gin"
)

/*收到请求设置trace日志*/
func AcceptRequestHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerMap = make(map[string]interface{})
		var queryMap = make(map[string]interface{})
		var paramMap = make(map[string]interface{})
		var bodyMap = make(map[string]interface{})

		for k := range c.Request.Header {
			headerMap[k] = c.GetHeader(k)
		}
		for k := range c.Request.URL.Query() {
			queryMap[k] = c.Query(k)
		}
		for _, param := range c.Params {
			paramMap[param.Key] = param.Value
		}

		if c.ContentType() == "application/json" {
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
				c.Request.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))
			}
			bodyJsonString := strings.Trim(string(bodyBytes), "\n")
			json.Unmarshal([]byte(bodyJsonString), &bodyMap)
		}

		type RequestDataType struct {
			Header map[string]interface{} `json:"header"`
			Query  map[string]interface{} `json:"query"`
			Body   map[string]interface{} `json:"body"`
			Param  map[string]interface{} `json:"param"`
		}

		requestData := RequestDataType{
			Header: headerMap,
			Query:  queryMap,
			Body:   bodyMap,
			Param:  paramMap,
		}
		requestDataJson, _ := json.MarshalIndent(requestData, "", "    ")
		logger.TraceLog.Info("[http-request] " + c.Request.Method + ": " + c.Request.URL.RequestURI() + " " + string(requestDataJson))

		c.Set("query", queryMap)
		c.Set("body", bodyMap)
		c.Set("param", paramMap)
		c.Next()
	}
}
