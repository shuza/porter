package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"user-service/model"
)

type UserRepository struct {
	conn *gorm.DB
}

func (repo *UserRepository) Init() error {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, DBName))
	if err != nil {
		return err
	}

	repo.conn = db
	return nil
}

func (repo *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := repo.conn.Find(&users).Error
	return users, err
}

func (repo *UserRepository) Get(id uint) (model.User, error) {
	var user model.User
	user.ID = id
	err := repo.conn.First(&user).Error
	return user, err
}

func (repo *UserRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	if err := repo.conn.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user interface{}) error {
	repo.conn.AutoMigrate(user)
	return repo.conn.Create(user).Error
}

func (repo *UserRepository) Close() {
	repo.conn.Close()
}
