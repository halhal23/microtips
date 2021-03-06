package main

import (
	"context"
	"log"
	"microtips/user/client"
	"microtips/user/pb"
)

func main() {
	c, err := client.NewClient("localhost:50052")
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}
	// create(c)
	update(c)
}

func create(c *client.Client) {
	input := &pb.UserInput{
		Name:     "kakiage",
		Password: "Shinobi",
	}
	req := &pb.CreateUserRequest{
		UserInput: input,
	}
	res, err := c.Service.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create user: %v\n", err)
	}
	log.Printf("user created successfully. id: %d, name: %s, password: %s\n", res.User.Id, res.User.Name, res.User.Password)
}

func update(c *client.Client) {
	input := &pb.UserInput{
		Name:     "hahahaha",
		Password: "hahahaha",
	}
	req := &pb.UpdateUserRequest{
		Id:        4,
		UserInput: input,
	}
	res, err := c.Service.UpdateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to update user: %v\n", err)
	}
	log.Printf("user updated successfully. id: %d, name: %s, password: %s\n", res.User.Id, res.User.Name, res.User.Password)
}
