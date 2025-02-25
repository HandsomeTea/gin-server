package env

import "os"

var defaultEnv = map[string]string{
	"GO_ENV": "development",
	"PORT":   "8084",
	// debug info warn error dpanic panic fatal
	"LOG_LEVEL":       "debug",
	"TRACE_LOG_LEVEL": "debug",
	// "DB_URL":          "mongodb://admin:admin@localhost:27017/gin?authSource=admin",
	"DB_URL": "mysql://root:root@0.0.0.0:3306/gin",
}

func GetEnv(key string) string {
	result := os.Getenv(key)

	if result != "" {
		return result
	}
	return defaultEnv[key]
}
