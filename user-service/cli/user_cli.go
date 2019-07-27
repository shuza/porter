package main

import (
	"context"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	pb "github.com/shuza/porter/user-service/proto"
	log "github.com/sirupsen/logrus"
)

func main() {
	registry := consul.NewRegistry()
	srv := k8s.NewService(
		micro.Name("porter.user-cli"),
		micro.Version("latest"),
		micro.Registry(registry),
	)
	srv.Init()

	client := pb.NewUserServiceClient("porter.auth", srv.Client())

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

	sng := make(chan int)
	<-sng
}

func getDummyUser() *pb.User {
	return &pb.User{
		Name:     "asd",
		Email:    "asd@github.com",
		Password: "123456",
		Company:  "BBC",
	}
}
