package database

import (
	"database/sql"
	"fmt"

	"github.com/RINOHeinrich/multiserviceauth/models"
)

// Structure pour gérer la connexion MySQL
type MySQL struct {
	config Dbconfig
	DB     *sql.DB
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
