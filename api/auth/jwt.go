package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(user googleUser) (string, error) {
	secret := os.Getenv("KEY")
	if secret == "" {
		return "", fmt.Errorf("KEY not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.Picture,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenHeader string) (*googleUser, error) {
	if !strings.HasPrefix(tokenHeader, "Bearer ") {
		return nil, fmt.Errorf("invalid authorization header")
	}

	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

	secret := os.Getenv("KEY")
	if secret == "" {
		return nil, fmt.Errorf("KEY not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &googleUser{
			ID:      claims["sub"].(string),
			Email:   claims["email"].(string),
			Name:    claims["name"].(string),
			Picture: claims["picture"].(string),
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
