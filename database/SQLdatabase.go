package database

import (
	"context"
	"fmt"

	"database/sql"

	"github.com/RINOHeinrich/multiserviceauth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Connect() error
	Disconnect() error
	Insert(data interface{}) error
	Update(id string, data interface{}) error
	Delete(id string) error
	Find(id string) (models.User, error)
	FindAll() ([]models.User, error)
	// Ajoutez d'autres méthodes selon vos besoins
}
type dbConfig struct {
	DBPort     int
	DBName     string
	DBHost     string
	DBPassword string
	DBUser     string
}

// Structure pour gérer la connexion MongoDB
type MongoDB struct {
	config dbConfig
	DB     *mongo.Client
}

// Structure pour gérer la connexion MySQL
type MySQL struct {
	config dbConfig
	DB     *sql.DB
}
type Postgres struct {
	config dbConfig
	DB     *sql.DB
}

func (m *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", m.config.DBUser, m.config.DBHost, m.config.DBPassword, m.config.DBPort))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	m.DB = client
	return nil
}

func (m *MongoDB) Disconnect() error {
	// Déconnexion de MongoDB
	err := m.DB.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoDB) Insert(data *models.User) error {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	_, err := userCollection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoDB) Delete(id string) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	_, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Printf("error when deleting users: \n: %s", err)
	}
}
func (m *MongoDB) Update(id string, data *models.User) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	update := bson.D{{Key: "$set", Value: data}}
	_, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("error when updating users: \n: %s", err)
	}
}
func (m *MongoDB) Find(id string) (*models.User, error) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	var user models.User
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *MongoDB) FindAll() ([]models.User, error) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
func (m *MySQL) Connect() error {
	// Connexion à MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.config.DBUser, m.config.DBPassword, m.config.DBHost, m.config.DBPort, m.config.DBName))
	if err != nil {
		return err
	}
	m.DB = db
	return nil
}
func (m *MySQL) Disconnect() error {
	// Déconnexion de MySQL
	err := m.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
func (m *MySQL) Insert(data *models.User) error {
	stmt, err := m.DB.Prepare("INSERT INTO users (id, Username, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.ID, data.Username, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil
}
func (m *MySQL) Delete(data *models.User) {
	stmt, err := m.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		fmt.Printf("error when deleting users: \n: %s", err)
	}
	_, err = stmt.Exec(data.ID)
	if err != nil {
		fmt.Printf("error when deleting users: \n: %s", err)
	}
}
func (m *MySQL) Update(id string, data *models.User) {
	stmt, err := m.DB.Prepare("UPDATE users SET Username = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Printf("error when updating users: \n: %s", err)
	}
	_, err = stmt.Exec(data.Username, data.Email, data.Password, id)
	if err != nil {
		fmt.Printf("error when updating users: \n: %s", err)
	}
}
func (m *MySQL) Find(id string) (*models.User, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *MySQL) FindAll() ([]models.User, error) {
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (m *Postgres) Connect() error {
	// Connexion à MySQL
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", m.config.DBHost, m.config.DBPort, m.config.DBUser, m.config.DBPassword, m.config.DBName))
	if err != nil {
		return err
	}
	m.DB = db
	return nil
}
func (m *Postgres) Disconnect() error {
	// Déconnexion de MySQL
	err := m.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
func (m *Postgres) Insert(data *models.User) error {
	stmt, err := m.DB.Prepare("INSERT INTO users (id, Username, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.ID, data.Username, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil
}
func (m *Postgres) Delete(id string) error {
	stmt, err := m.DB.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func (m *Postgres) Update(id string, data *models.User) error {
	stmt, err := m.DB.Prepare("UPDATE users SET Username = $1, email = $2, password = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Username, data.Email, data.Password, id)
	if err != nil {
		return err
	}
	return nil
}
func (m *Postgres) Find(id string) (*models.User, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *Postgres) FindAll() ([]models.User, error) {
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func Find(db Database, id string) (*models.User, error) {
	user, err := db.Find(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
	// Implémentez la méthode Find pour MongoDB
}
func FindAll(db Database) ([]models.User, error) {
	users, err := db.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
	// Implémentez la méthode FindAll pour MongoDB
}
func Insert(db Database, data *models.User) error {
	err := db.Insert(data)
	if err != nil {
		return err
	}
	return nil
	// Implémentez la méthode Insert pour MongoDB
}
func Update(db Database, id string, data *models.User) error {
	err := db.Update(id, data)
	if err != nil {
		return err
	}
	return nil
	// Implémentez la méthode Update pour MongoDB
}
func Delete(db Database, id string) error {
	err := db.Delete(id)
	if err != nil {
		return err
	}
	return nil
	// Implémentez la méthode Delete pour MongoDB
}

// Implémentez les autres méthodes de l'interface Database pour MongoDB
