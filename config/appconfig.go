package config

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppDatabases struct {
	MongoDB *mongo.Client
}

func SetAppConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
}

func InitDatabases() *AppDatabases {
	mongo := initMongoDB()
	return &AppDatabases{MongoDB: mongo}
}

func (dbs *AppDatabases) DeferDatabases() {
	dbs.MongoDB.Disconnect(nil)
}

func initMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongodb.uri")))
	if err != nil {
		fmt.Println("MongoDB error connection!")
	}
	return client
}
