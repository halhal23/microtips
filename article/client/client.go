package client

import (
	"microtips/article/pb"

	"google.golang.org/grpc"
)

// client から　serviceを利用できるようにする
type Client struct {
	conn    *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	// client connection を生成
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// articleサービスの　client を生成
	c := pb.NewArticleServiceClient(conn)

	// articleサービスのクライアントを返す
	return &Client{conn, c}, nil
}
