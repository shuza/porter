package main

import (
	"github.com/shuza/porter/user-service/db"
	pb "github.com/shuza/porter/user-service/proto"
	"github.com/shuza/porter/user-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	dbConn, err := db.CreateDb()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	dbConn.AutoMigrate(&pb.User{})
	repo := &db.UserRepository{dbConn}
	tokenService := &service.TokenService{repo}

	port := os.Getenv("PORT")

	//	setup gRPC server
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen  :  ", err)
	}

	s := grpc.NewServer()
	userService := service.UserService{repo, tokenService}
	pb.RegisterUserServiceServer(s, &userService)
	reflection.Register(s)

	log.Println("Running on port :  ", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to server  :  ", err)
	}

}
