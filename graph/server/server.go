package main

import (
	"log"
	"microtips/article/client"
	"microtips/graph"
	"microtips/graph/generated"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up Gin
	r := gin.Default()
	///////////////////
	// CORS
	///////////////////
	r.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
		// AllowOriginFunc: func(origin string) bool {
		//  return origin == "https://www.example.com:8080"
		// },
		// preflight requestで許可した後の接続可能時間
		// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
		MaxAge: 24 * time.Hour,
	}))
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// articleClientを生成
	articleClient, err := client.NewClient("localhost:50051")
	if err != nil {
		articleClient.Close()
		log.Fatalf("Failed to create article client: %v\n", err)
	}
	// h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{ArticleClient: articleClient}}))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	// h := handler.Playground("GraphQL", "/query")
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
