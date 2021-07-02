package main

import (
	"log"
	"microtips/user/pb"
	"microtips/user/repository"
	"microtips/user/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer listener.Close()

	repo, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create repository: %v\n", err)
	}

	svc := service.NewService(repo)

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, svc)

	log.Println("listening to 50052...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
