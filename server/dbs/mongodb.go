package dbs

import (
	"context"
	env "gin-server/server/configs/env"
	logger "gin-server/server/configs/logger"
	"net/url"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Mongodb *mongo.Database

func ConnectMongodb() {
	mongodbAddress := env.GetEnv("DB_URL")
	client, err := mongo.Connect(options.Client().ApplyURI(mongodbAddress))

	if err != nil {
		logger.SystemLog.Error("connect mongodb: " + mongodbAddress + " failed: " + err.Error())
		client.Disconnect(context.TODO())
		panic(err)
	}
	logger.SystemLog.Info("mongodb connected on " + mongodbAddress + " success and ready to use.")

	dbUrl, _ := url.Parse(mongodbAddress)
	dbName := dbUrl.Path[1:]

	Mongodb = client.Database(dbName)
}
