package main

import (
	"github.com/micro/go-micro"
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
	"github.com/shuza/porter/user-service/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	dbConn, err := db.CreateDb()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	dbConn.AutoMigrate(&pb.User{})

	repo := db.NewUserRepository(dbConn)
	tokenService := service.NewTokenService(&repo)

	srv := micro.NewService(
		micro.Name("porter.auth"),
	)
	srv.Init()

	userService := service.NewUserService(&repo, &tokenService)
	pb.RegisterUserServiceHandler(srv.Server(), &userService)

	log.Println("Auth service is running...")
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
