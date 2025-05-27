package mongo

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	mongo      *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func NewClient(ctx context.Context, connectionString, database, collection string) (*MongoClient, error) {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	col := db.Collection(collection)

	return &MongoClient{mongo: client, db: db, collection: col}, nil
}

func (c *MongoClient) GetCollection() *mongo.Collection {
	return c.collection
}

func (c *MongoClient) Close(ctx context.Context) {
	if err := c.mongo.Disconnect(ctx); err != nil {
		slog.ErrorContext(ctx, "Error when trying disconnect from mongo", slog.Any("error", err))
	}
}
