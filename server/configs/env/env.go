package env

import (
	"os"
	"sync"
)

var defaultEnv *sync.Map

func init() {
	defaultEnv = &sync.Map{}
	defaultEnv.Store("GO_ENV", "development")
	defaultEnv.Store("PORT", "8084")
	defaultEnv.Store("LOG_LEVEL", "debug") // log level value: debug info warn error dpanic panic fatal
	defaultEnv.Store("TRACE_LOG_LEVEL", "debug")
	defaultEnv.Store("OTEL_ENABLED", "no")
	defaultEnv.Store("DB_URL", "mysql://root:root@0.0.0.0:3306/gin") // mongodb://admin:admin@localhost:27017/gin?authSource=admin
}

func GetEnv(key string) (string, bool) {
	if val, ok := os.LookupEnv(key); ok {
		return val, true
	}
	if val, ok := defaultEnv.Load(key); ok {
		return val.(string), true
	}
	return "", false
}
