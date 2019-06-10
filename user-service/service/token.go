package service

import "github.com/shuza/porter/user-service/db"

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService struct {
	repo db.IRepository
}

func (s *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

func (s *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
