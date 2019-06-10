package service

import (
	"context"
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
)

type UserService struct {
	repo         db.IRepository
	tokenService Authable
}

func (s *UserService) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: user}, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &pb.Response{Users: users}, nil
}

func (s *UserService) Auth(ctx context.Context, req *pb.User) (*pb.Token, error) {
	_, err := s.repo.GetByEmailAndPassword(req)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: "testingabc"}, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	return &pb.Token{}, nil
}
