package user

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) GetUserByMobile(mobile string) (*model.User, error) {
	userM, err := u.userRepo.GetByMobile(mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrUsertNotFound
		}
		return nil, err
	}
	return userM, nil
}

func (u *userService) GetUserDetailsByMobile(mobile string) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	userM, err := u.userRepo.GetByMobile(mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared.ErrUsertNotFound,
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
	userDetails := user.ResponseDetails{
		ID:           userM.ID,
		PhoneNumber:  userM.PhoneNumber,
		FirstName:    userM.FirstName,
		LastName:     userM.LastName,
		Email:        userM.Email,
		RegisteredAt: userM.CreatedAt,
		HasPassword:  userM.Password != "",
	}
	return &userDetails, nil
}
