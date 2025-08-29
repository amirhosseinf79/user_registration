package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type authHandler struct {
	authService interfaces.AuthService1
}

func NewAuthHandler(authService interfaces.AuthService1) interfaces.AuthHandler1 {
	return &authHandler{
		authService: authService,
	}
}
