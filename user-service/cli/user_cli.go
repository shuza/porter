package main

import (
	"context"
	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	pb "github.com/shuza/porter/user-service/proto"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)
	srv.Init()

	client := pb.NewUserServiceClient("porter.auth", microClient.DefaultClient)

	r, err := client.Create(context.TODO(), getDummyUser())
	if err != nil {
		log.Fatalf("Could not create user  Error  :   %v\n", err)
	}
	log.Printf("Created User ID :  %v\n", r.User.Id)

	getAll, err := client.GetAll(context.TODO(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not list user  Error  :   %v\n", err)
	}

	//	Print user list
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), getDummyUser())
	if err != nil {
		log.Fatalf("Could not authenticate user  Error  :   %v\n", err)
	}

	log.Printf("User token is  %v\n", authResponse.Token)

	os.Exit(0)
}

func getDummyUser() *pb.User {
	return &pb.User{
		Name:     "asd",
		Email:    "asd@github.com",
		Password: "123456",
		Company:  "BBC",
	}
}
