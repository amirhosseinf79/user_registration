package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
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

func (j *jwtService) GenerateAuthTokens(userID uint) (*auth.ResponseJWT, *shared_dto.ResponseOneMessage) {
	accessToken, err := j.jwtRepo.GenerateToken(userID, false)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	refreshToken, err := j.jwtRepo.GenerateToken(userID, true)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	tokenM := model.Token{
		RefreshToken: refreshToken,
		UserID:       userID,
	}
	err = j.tokenRepo.SaveRefreshToken(&tokenM)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	token := auth.ResponseJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &token, nil
}

func (j *jwtService) GetUserIDByRefreshToken(oldRefreshToken string) (uint, *shared_dto.ResponseOneMessage) {
	userID, err := j.tokenRepo.GetUserIDByRefresh(oldRefreshToken)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return 0, result
	}
	return userID, nil
}

func (j *jwtService) GetUserIDByAccessToken(accessToke string) (uint, *shared_dto.ResponseOneMessage) {
	userID, err := j.jwtRepo.Verify(accessToke)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return 0, result
	}
	return userID, nil
}
