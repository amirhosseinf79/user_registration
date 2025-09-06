package user

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) UpdateUserPassword(userID uint, fields user.FieldUpdatePassword) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	userM, err := u.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared.ErrUserNotFound,
				RealError:  err,
			})
			return nil, result
		}
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	if userM.Password != "" && !userM.ValidatePassword(fields.OldPassword) {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidPassword,
		})
		return nil, result
	}

	hashedPassword, err := pkg.HashPassword(fields.NewPassword)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	userM.Password = hashedPassword
	err = u.userRepo.Update(userM)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userDetails := user.NewUserResponse(userM)
	return userDetails, nil
}
