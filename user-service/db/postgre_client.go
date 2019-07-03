package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	pb "github.com/shuza/porter/user-service/proto"
	"os"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return UserRepository{conn}
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

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	err := repo.conn.Find(&users).Error
	return users, err
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	err := repo.conn.First(&user).Error
	return user, err
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	repo.conn.AutoMigrate(user)
	return repo.conn.Create(user).Error
}

func (repo *UserRepository) Close() {
	repo.conn.Close()
}
