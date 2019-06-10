package service

import (
	"context"
	"errors"
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	log.Printf("Logging in with :  %v  %v\n", req.Email, req.Password)
	user, err := s.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return nil, err
	}

	//	Compare our given password with the hashed password stored in database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := s.tokenService.Encode(user)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	claim, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	if claim.User.Id == "" {
		return nil, errors.New("invalid user")
	}

	return &pb.Token{Valid: true}, nil
}
