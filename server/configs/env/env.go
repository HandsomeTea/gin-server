package env

import "os"

var defaultEnv = map[string]string{
	"GO_ENV":          "development",
	"PORT":            "8084",
	"LOG_LEVEL":       "debug", // log level value: debug info warn error dpanic panic fatal
	"TRACE_LOG_LEVEL": "debug",
	"DB_URL":          "mysql://root:root@0.0.0.0:3306/gin", // mongodb://admin:admin@localhost:27017/gin?authSource=admin
}

func GetEnv(key string) string {
	result := os.Getenv(key)

	if result != "" {
		return result
	}
	return defaultEnv[key]
}
