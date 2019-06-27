package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
	"time"
)

var (
	//	Define a secure key string used
	//	as salt when hashing our tokens
	key = []byte("hashingpasswordismandatory")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo db.IRepository
}

func (s *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	//	Parse token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	//	Validate the token and return custom claim
	if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claim, nil
	}

	return nil, err
}

//	Encode claim into JWT
func (s *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(72 * time.Hour).Unix()

	//	Create claim
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "porter.user",
		},
	}

	//	Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//	Sign token and return
	return token.SignedString(key)

}
