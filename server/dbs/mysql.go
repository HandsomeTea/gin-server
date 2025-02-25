package dbs

import (
	env "gin-server/server/configs/env"
	logger "gin-server/server/configs/logger"
	"net/url"

	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func getDbDSN() (string, string) {
	mysqlAddress := env.GetEnv("DB_URL")
	dbUrl, _ := url.Parse(mysqlAddress)

	userName := dbUrl.User.Username()
	password, _ := dbUrl.User.Password()
	dbName := dbUrl.Path[1:]
	queryStr := dbUrl.RawQuery

	if queryStr != "" {
		queryStr = "?" + queryStr
	}
	return userName + ":" + password + "@tcp(" + dbUrl.Host + ")/" + dbName + queryStr, mysqlAddress
}

// var Mysql *sql.DB

// // go-sql-driver连接数据库
// func ConnectMysql() {
// 	dsn, mysqlAddress := getDbDSN()
// 	db, _err := sql.Open("mysql", dsn)

// 	if _err != nil {
// 		logger.SystemLog.Error("connect mysql: " + mysqlAddress + " failed: " + _err.Error())
// 		panic(_err)
// 	}
// 	Mysql = db
// 	err := Mysql.Ping()

// 	if err != nil {
// 		logger.SystemLog.Error("connect mysql: " + mysqlAddress + " failed: " + err.Error())
// 		panic(err)
// 	}
// 	logger.SystemLog.Info("mysql connected on " + mysqlAddress + " success and ready to use.")
// }

var Mysql *gorm.DB

// gorm连接数据库
func ConnectMysql() {
	dsn, mysqlAddress := getDbDSN()
	db, _err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		},
	})

	if _err != nil {
		logger.SystemLog.Error("connect mysql: " + mysqlAddress + " failed: " + _err.Error())
		panic(_err)
	}
	Mysql = db
	logger.SystemLog.Info("mysql connected on " + mysqlAddress + " success and ready to use.")
}
