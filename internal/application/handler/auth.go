package handler

import "github.com/amirhosseinf79/user_registration/internal/domain/interfaces"

type authHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) interfaces.AuthHandler {
	return &authHandler{
		authService: authService,
	}
}
