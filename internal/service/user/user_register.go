package user

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

func (u *userService) RegisterUserByNumber(phoneNumber string) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	exists, err := u.userRepo.CheckMobileExists(phoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	userM := &model.User{
		PhoneNumber: phoneNumber,
	}
	if !exists {
		err = u.userRepo.Create(userM)
		if err != nil {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
	} else {
		userM, err = u.userRepo.GetByMobile(phoneNumber)
		if err != nil {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
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
