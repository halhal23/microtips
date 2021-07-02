package auth

import (
	"context"
	"log"
	"microtips/user/pb"
	"microtips/user/pkg/jwt"
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
		log.Println("Authorization Header:", header)

		// Allow unauthenticated users in
		if header == "" {
			log.Println("Non Header:", header)
			return
		}

		//validate jwt token
		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		log.Println("Parsed token and get username:", username)
		if err != nil {
			return
		}
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *pb.User {
	raw, _ := ctx.Value(userCtxKey).(*pb.User)
	return raw
}
