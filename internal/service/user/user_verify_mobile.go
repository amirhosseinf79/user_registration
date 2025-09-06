package user

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (u *userService) VerifyUserMobile(userID uint, code string) (*user.ResponseDetails, *shared.ResponseOneMessage) {
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

	if userM.PhoneNumber == "" {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidMobile,
		})
		return nil, result
	}

	_, err2 := u.otpService.CheckOTPCode(otp.FieldVerifyOTP{
		FieldOTPStore: otp.FieldOTPStore{
			Prefix: "verify:mobile:",
			Key:    userM.PhoneNumber,
		},
		Code: code,
	})
	if err2 != nil {
		return nil, err2
	}

	userM.MobileVerified = true

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
