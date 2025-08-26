package jwt

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (j *jwtService) GetUserIDByRefreshToken(oldRefreshToken string) (uint, *shared.ResponseOneMessage) {
	userID, err := j.tokenRepo.GetUserIDByRefresh(oldRefreshToken)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return 0, result
	}
	return userID, nil
}

func (j *jwtService) GetUserIDByAccessToken(accessToke string) (uint, *shared.ResponseOneMessage) {
	userID, err := j.jwtRepo.Verify(accessToke)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return 0, result
	}
	return userID, nil
}
