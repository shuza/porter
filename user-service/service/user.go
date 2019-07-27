package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	repo         db.IRepository
	tokenService Authable
}

func NewUserService(repo db.IRepository, tokenService Authable) UserService {
	return UserService{repo: repo, tokenService: tokenService}
}

func (s *UserService) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	log.Println("starting create")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := s.repo.Create(req); err != nil {
		return errors.New(fmt.Sprintf("Create user DB Error :  %v\n", err))
	}
	resp.User = req
	token, err := s.tokenService.Encode(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Create user token create  Error  :   %v\n", err))
	}
	resp.User = req
	resp.Token = &pb.Token{Token: token}

	return nil
}

func (s *UserService) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = user

	return nil
}

func (s *UserService) GetAll(ctx context.Context, req *pb.Empty, resp *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users

	return nil
}

func (s *UserService) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	log.Printf("Logging in with :  %v  %v\n", req.Email, req.Password)
	user, err := s.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	//	Compare our given password with the hashed password stored in database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := s.tokenService.Encode(user)
	if err != nil {
		return err
	}
	resp.Token = token

	return nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	claim, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claim.User.Id == "" {
		return errors.New("invalid user")
	}
	resp.Valid = true

	return nil
}
