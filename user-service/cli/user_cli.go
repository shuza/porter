package main

import (
	"context"
	pb "github.com/shuza/porter/user-service/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

func amain() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Can't connect to server  :  %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, err := client.Create(context.TODO(), &pb.User{
		Name:     "asd",
		Email:    "asd@github.com",
		Password: "123456",
		Company:  "BBC",
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", resp.User.Id)

	respAll, err := client.GetAll(context.Background(), &pb.Empty{})
	for _, v := range respAll.Users {
		log.Println(v)
	}

	log.Println("\n\t=========	Auth	========")
	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    "asd@github.com",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: asd@github.com error: %v\n", err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)
	os.Exit(1)
}
