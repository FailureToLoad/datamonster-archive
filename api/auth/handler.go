package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/failuretoload/datamonster/request"
	"github.com/failuretoload/datamonster/response"
)

func GoogleAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(ctx, w, errors.New("missing or invalid authorization header"))
			return
		}

		user, err := validateJWT(authHeader)
		if err != nil {
			response.Unauthorized(ctx, w, err)
			return
		}

		ctx = context.WithValue(ctx, request.UserIDKey, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
