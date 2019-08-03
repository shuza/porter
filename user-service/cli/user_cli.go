package main

import (
	"context"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/registry/kubernetes"
	pb "github.com/shuza/porter/user-service/proto"
	log "github.com/sirupsen/logrus"
)

func main() {
	registry := kubernetes.NewRegistry()
	srv := k8s.NewService(
		micro.Name("porter.user-cli"),
		micro.Version("latest"),
		micro.Registry(registry),
	)
	srv.Init()

	if list, err := registry.ListServices(); err != nil {

	} else {
		log.Warnf("")
	}

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

	sng := make(chan int)
	<-sng
}

func createPacket(token string) {

}

func getDummyUser() *pb.User {
	return &pb.User{
		Name:     "user-1",
		Email:    "admin@github.com",
		Password: "123456",
		Company:  "BBC",
	}
}
