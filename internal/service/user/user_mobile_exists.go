package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (u *userService) CheckUserMobileExists(mobile string) *shared.ResponseOneMessage {
	exists, err := u.userRepo.CheckMobileExists(mobile)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return result
	}
	if exists {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusConflict,
			ErrMessage: shared.ErrMobileExists,
			RealError:  err,
		})
		return result
	}
	return nil
}
