package main

import (
	"fmt"
	"user-service/api"
	"user-service/db"
)

func main() {
	initDB()
	defer db.Client.Close()

	r := api.NewGinEngine()
	fmt.Println("Box service is running on port 8083 ....")
	if err := r.Run(":8083"); err != nil {
		panic(err)
	}

	/*repo := db.NewUserRepository(dbConn)
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
	}*/
}

func initDB() {
	db.Client = &db.UserRepository{}
	if err := db.Client.Init(); err != nil {
		panic(err)
	}
}
