package databaseNoSQL

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSQLDatabaseConnector interface {
	Connect() (*mongo.Client, error)
}
type MongoDBConnector struct {
	// MongoDB specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Document string
}

func (m *MongoDBConnector) Connect() (*mongo.Client, error) {
	// Implement the connection logic for MongoDB using mongo-driver
	// Example:
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", m.Username, m.Password, m.Host, m.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	return client, err
}

type DatabaseNoSQL struct {
	Connector NoSQLDatabaseConnector
}
