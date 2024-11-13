package main

import (
	pb "auth_service/pkg/api/auth"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewAuthServiceClient(conn)

	// /RegisterUser
	{
		_, err := cli.Register(context.Background(), &pb.RegisterRequest{
			Email:    "email",
			Password: "password",
		})
		if err != nil {
			log.Fatalf("RegisterUser error: %v", err)
		} else {
			log.Printf("user success register")
		}
	}

	// /LoginUser
	{
		resp, err := cli.Login(context.Background(), &pb.LoginRequest{
			Email:    "email",
			Password: "password",
		})
		if err != nil {
			log.Fatalf("LoginErr error: %v", err)
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			notes, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf(" protojson.Marshal error: %v", err)
			} else {
				log.Printf("notes: %s", string(notes))
			}
		}
	}
}
