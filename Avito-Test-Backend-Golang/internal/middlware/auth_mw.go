package middleware

import (
	"Avito-Test-Backend-Golang/internal/util/jwter"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func New(log *logrus.Logger, jwtManager *jwter.JwtManager) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("start auth middleware")
			header := r.Header
			value := header.Get("Authorization")
			jwtToken := strings.Replace(value, "Bearer ", "", 1)
			if jwtToken == "" {
				log.Error("jwt token is empty")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, err := jwtManager.DecodeToken(jwtToken)
			if err != nil {
				log.Error("jwt token is invalid")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "userId", claims.Principal)
			ctx = context.WithValue(ctx, "is_admin", claims.IsAdmin)
			ctx = context.WithValue(ctx, "tagId", claims.TagId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
