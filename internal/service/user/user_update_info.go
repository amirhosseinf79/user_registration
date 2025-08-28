package user

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) UpdateUserProfile(userID uint, fields user.FieldUpdateDetails) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	userM, err := u.userRepo.GetByID(userID)
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

	if fields.FirstName != "" {
		userM.FirstName = fields.FirstName
	}
	if fields.LastName != "" {
		userM.LastName = fields.LastName
	}
	if fields.Email != "" {
		err := u.CheckUserEmailExists(fields.Email)
		if err != nil {
			return nil, err
		}
		userM.Email = fields.Email
		userM.EmailVerified = false
	}
	if fields.PhoneNumber != "" {
		err := u.CheckUserMobileExists(fields.PhoneNumber)
		if err != nil {
			return nil, err
		}
		userM.PhoneNumber = fields.PhoneNumber
		userM.MobileVerified = false
	}

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
