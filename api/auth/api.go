package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/failuretoload/datamonster/response"
	"google.golang.org/api/idtoken"
)

type googleAuthRequest struct {
	Credential string `json:"credential"`
}

type googleUser struct {
	ID      string `json:"userId"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) HandleGoogleAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req googleAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(ctx, w, "invalid request body", err)
		return
	}

	payload, err := verifyGoogleToken(ctx, req.Credential)
	if err != nil {
		response.Unauthorized(ctx, w, err)
		return
	}

	user := googleUser{
		ID:      payload.Claims["sub"].(string),
		Email:   payload.Claims["email"].(string),
		Name:    payload.Claims["name"].(string),
		Picture: payload.Claims["picture"].(string),
	}

	token, err := createJWT(user)
	if err != nil {
		response.InternalServerError(ctx, w, "failed to create token", err)
		return
	}

	response.OK(ctx, w, map[string]interface{}{
		"token":   token,
		"userId":  user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.Picture,
	})
}

func (c Controller) ValidateToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("Authorization")
	if token == "" {
		response.Unauthorized(ctx, w, errors.New("missing authorization header"))
		return
	}

	user, err := validateJWT(token)
	if err != nil {
		response.Unauthorized(ctx, w, err)
		return
	}

	response.OK(ctx, w, user)
}

func verifyGoogleToken(ctx context.Context, token string) (*idtoken.Payload, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("GOOGLE_CLIENT_ID not set")
	}

	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	return payload, nil
}
