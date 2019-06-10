package db

import (
	"github.com/jinzhu/gorm"
	pb "github.com/shuza/porter/user-service/proto"
)

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
