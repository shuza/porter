package db

import (
	pb "github.com/shuza/porter/user-service/proto"
)

type IRepository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
	Create(user *pb.User) error
}
