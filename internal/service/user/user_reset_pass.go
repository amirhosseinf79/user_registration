package user

import (
	"errors"
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) ResetUserPasswordByInfo(userInfo, newPassword string) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	var err error
	var userM *model.User
	reg := regexp.MustCompile(`^09\d{9}$`)
	if reg.MatchString(userInfo) {
		userM, err = u.userRepo.GetByMobile(userInfo)
	} else {
		userM, err = u.userRepo.GetByEmail(userInfo)
	}
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

	hashedPassword, err := pkg.HashPassword(newPassword)
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
