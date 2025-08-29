package auth_middleware

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type authMiddleware struct {
	jwtService interfaces.JWTService
	prefix     string
}

func NewAuthMiddleware(jwtService interfaces.JWTService) interfaces.AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
		prefix:     "Bearer ",
	}
}
