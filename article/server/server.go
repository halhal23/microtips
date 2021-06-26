package server

import (
	"log"
	"microtips/article/pb"
	"microtips/article/repository"
	"microtips/article/service"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// article サーバに接続
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// repositoryを作成
	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	// Serviceを作成
	service := service.NewService(repository)

	// サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	// サーバーを起動
	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
