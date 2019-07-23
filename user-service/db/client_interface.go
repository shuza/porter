package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	pb "github.com/shuza/porter/user-service/proto"
	"os"
)

type IRepository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
	Create(user *pb.User) error
}

func CreateDb() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	return gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, DBName))
}
