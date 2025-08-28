package user

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (u *userService) RegisterUserByEmail(fields auth.FieldEmailRegister) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	err := u.CheckUserEmailExists(fields.Email)
	if err != nil {
		return nil, err
	}

	hashPassword, err2 := pkg.HashPassword(fields.Password)
	if err2 != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err2,
		})
		return nil, result
	}

	userM := &model.User{
		FirstName:   fields.FirstName,
		LastName:    fields.LastName,
		PhoneNumber: fields.PhoneNumber,
		Email:       fields.Email,
		Password:    hashPassword,
	}
	err2 = u.userRepo.Create(userM)
	if err2 != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err2,
		})
		return nil, result
	}

	userDetails := user.NewUserResponse(userM)
	return userDetails, nil
}
