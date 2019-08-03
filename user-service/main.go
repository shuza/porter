package main

import (
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
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

	registry := kubernetes.NewRegistry()
	srv := k8s.NewService(
		micro.Name("porter.auth"),
		micro.Version("latest"),
		micro.Registry(registry),
	)
	srv.Init()

	userService := service.NewUserService(&repo, &tokenService)
	pb.RegisterUserServiceHandler(srv.Server(), &userService)

	log.Println("Auth service is running...")
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
