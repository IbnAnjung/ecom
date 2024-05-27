package nosql

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	User     string
	Password string
	Host     string
	Source   string
}

type Mongo struct {
	Client *mongo.Client
}

func NewMongoConnection(ctx context.Context, conf MongoConfig) (m Mongo, err error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", conf.User, conf.Password, conf.Host, conf.Source)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	loggerOptions := options.Logger().SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetLoggerOptions(loggerOptions)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return
	}

	if err = client.Ping(ctx, nil); err != nil {
		return
	}

	return Mongo{
		Client: client,
	}, nil

}

func (m Mongo) Cleanup(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
