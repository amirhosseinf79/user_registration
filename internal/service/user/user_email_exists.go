package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (u *userService) CheckUserEmailExists(email string) *shared.ResponseOneMessage {
	exists, err := u.userRepo.CheckEmailExists(email)
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
			ErrMessage: shared.ErrEmailExists,
			RealError:  err,
		})
		return result
	}
	return nil
}
