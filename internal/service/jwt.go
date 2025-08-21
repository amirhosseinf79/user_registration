package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
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

func (j *jwtService) GenerateAuthTokens(userID uint) (*dto.ResponseAuthOk, error) {
	accessToken, err := j.jwtRepo.GenerateToken(userID, false)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.jwtRepo.GenerateToken(userID, true)
	if err != nil {
		return nil, err
	}
	tokenM := model.Token{
		RefreshToken: refreshToken,
		UserID:       userID,
	}
	err = j.tokenRepo.SaveRefreshToken(&tokenM)
	if err != nil {
		return nil, err
	}

	token := dto.ResponseAuthOk{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &token, nil
}

func (j *jwtService) GetUserIDByRefreshToken(oldRefreshToken string) (uint, error) {
	userID, err := j.tokenRepo.GetUserIDByRefresh(oldRefreshToken)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
