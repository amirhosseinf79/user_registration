package user

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) GetUserByID(id uint) (*model.User, error) {
	userM, err := u.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrUserNotFound
		}
		return nil, err
	}
	return userM, nil
}

func (u *userService) GetUserDetailsByID(userID uint) (*user.ResponseDetails, *shared.ResponseOneMessage) {
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
	userDetails := user.NewUserResponse(userM)
	return userDetails, nil
}
