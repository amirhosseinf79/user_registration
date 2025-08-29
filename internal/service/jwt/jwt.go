package jwt

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type jwtService struct {
	jwtRepo   repository.JWTRepository
	tokenRepo repository.TokenRepository
}

func NewJWTService(jwtRepo repository.JWTRepository, tokenRepo repository.TokenRepository) interfaces.JWTService {
	return &jwtService{
		jwtRepo:   jwtRepo,
		tokenRepo: tokenRepo,
	}
}
