package auth

import (
	"context"
	"log"
	"microtips/user/pb"
	"microtips/user/pkg/jwt"
	"microtips/user/repository"
	"microtips/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func LoginMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			// username, err := jwt.ParseToken(tokenStr)
			_, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		// Allow unauthenticated users in
		if header == "" {
			log.Println("Non Header:", header)
			return
		}

		//validate jwt token
		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			return
		}

		repo, err := repository.NewsqliteRepo()
		if err != nil {
			log.Printf("Cannot initialize repository: %v\n", err)
			return
		}
		svc := service.NewService(repo)
		res, err := svc.ReadUserByName(c, &pb.ReadUserByNameRequest{Name: username})
		if err != nil {
			log.Printf("Cannot find user by name: %v\n", err)
			return
		}
		log.Printf("finded user in auth middleware: %v\n", res.User)

		// and call the next with our new context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, &pb.User{Id: res.User.Id, Name: res.User.Name, Password: res.User.Password})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *pb.User {
	raw, _ := ctx.Value(userCtxKey).(*pb.User)
	return raw
}
