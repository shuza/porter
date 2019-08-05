package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user-service/model"
)

type IRepository interface {
	Init() error
	GetAll() ([]model.User, error)
	Get(id uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Create(user interface{}) error
	Close()
}

var Client IRepository
