package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/RINOHeinrich/multiserviceauth/database/databaseNoSQL"
	"github.com/RINOHeinrich/multiserviceauth/database/databaseSQL"
	"github.com/joho/godotenv"
)

type Database struct {
	SQLDB   *sql.DB
	NoSQLDB *mongo.Client
}

var DB Database

func init() {
	godotenv.Load("config/.env")
	dbType := os.Getenv("DB_TYPE")
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error converting port to int: ", err)
	}
	dbname := os.Getenv("DB_NAME")

	switch dbType {
	case "mysql":
		DatabaseConnection := databaseSQL.DatabaseSQL{Connector: &databaseSQL.MySQLConnector{
			Host:     dbhost,
			Port:     dbport,
			Username: dbuser,
			Password: dbpassword,
			Database: dbname,
		}}
		DB.SQLDB, err = DatabaseConnection.Connector.Connect()
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
		}
	case "postgresql":
		DatabaseConnection := databaseSQL.DatabaseSQL{Connector: &databaseSQL.PostgreSQLConnector{
			Host:     dbhost,
			Port:     dbport,
			Username: dbuser,
			Password: dbpassword,
			Database: dbname,
		}}
		DB.SQLDB, err = DatabaseConnection.Connector.Connect()
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
		}
		fmt.Println("Connected to database: ", DB)
	case "mongodb":
		DatabaseConnection := databaseNoSQL.DatabaseNoSQL{Connector: &databaseNoSQL.MongoDBConnector{
			Host:     dbhost,
			Port:     dbport,
			Username: dbuser,
			Password: dbpassword,
			Document: dbname,
		}}
		DB.NoSQLDB, err = DatabaseConnection.Connector.Connect()
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
		}
		fmt.Println("Connected to database: ", DB)
	default:
		fmt.Println("Database type not supported")
	}
}
