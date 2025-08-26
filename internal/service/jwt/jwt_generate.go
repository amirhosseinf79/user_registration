package jwt

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (j *jwtService) GenerateAuthTokens(userID uint) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	accessToken, err := j.jwtRepo.GenerateToken(userID, false)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	refreshToken, err := j.jwtRepo.GenerateToken(userID, true)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
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
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
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
