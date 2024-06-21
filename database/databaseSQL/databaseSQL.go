package databaseSQL

import (
	"fmt"

	"database/sql"
)

type SQLDatabaseConnector interface {
	Connect() (*sql.DB, error)
}

func Connect(connector SQLDatabaseConnector) (*sql.DB, error) {
	db, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type MySQLConnector struct {
	// MySQL specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (m *MySQLConnector) Connect() (*sql.DB, error) {
	// Implement the connection logic for MySQL using gorm
	// Example:
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	//open connection to mysql using database/sql package
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	//ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type PostgreSQLConnector struct {
	// PostgreSQL specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (p *PostgreSQLConnector) Connect() (*sql.DB, error) {
	// Implement the connection logic for PostgreSQL using gorm
	// Example:
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
	//open connection to postgres using database/sql package
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	//ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type DatabaseSQL struct {
	Connector SQLDatabaseConnector
}
