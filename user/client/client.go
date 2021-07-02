package client

import (
	"microtips/user/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	Service pb.UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)
	return &Client{conn, client}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
