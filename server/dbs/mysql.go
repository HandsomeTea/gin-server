package dbs

import (
	"database/sql"
	env "gin-server/server/configs/env"
	logger "gin-server/server/configs/logger"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

var Mysql *sql.DB

func ConnectMysql() {
	mysqlAddress := env.GetEnv("DB_URL")
	dbUrl, _ := url.Parse(mysqlAddress)

	userName := dbUrl.User.Username()
	password, _ := dbUrl.User.Password()
	dbName := dbUrl.Path[1:]
	queryStr := dbUrl.RawQuery

	if queryStr != "" {
		queryStr = "?" + queryStr
	}

	db, _err := sql.Open("mysql", userName+":"+password+"@tcp("+dbUrl.Host+")/"+dbName+queryStr)

	if _err != nil {
		logger.SystemLog.Error("connect mysql: " + mysqlAddress + " failed: " + _err.Error())
		panic(_err)
	}
	Mysql = db
	err := Mysql.Ping()

	if err != nil {
		logger.SystemLog.Error("connect mysql: " + mysqlAddress + " failed: " + err.Error())
		panic(err)
	}
	logger.SystemLog.Info("mysql connected on " + mysqlAddress + " success and ready to use.")
}
