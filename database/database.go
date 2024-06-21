package database

import "github.com/RINOHeinrich/multiserviceauth/models"

type Database interface {
	Connect() error
	Disconnect() error
	Insert(data *models.User) error
	Update(id string, data *models.User) error
	Delete(id string) error
	Find(id string) (*models.User, error)
	FindAll() ([]models.User, error)
	// Ajoutez d'autres méthodes selon vos besoins
}
type Dbconfig struct {
	DBPort     int
	DBName     string
	DBHost     string
	DBPassword string
	DBUser     string
}

func Find(db Database, id string) (*models.User, error) {
	user, err := db.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
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
