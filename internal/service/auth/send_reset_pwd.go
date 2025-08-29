package auth

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
)

func (a *authService) SendResetPasswerd(fields auth.FieldSendResetPwd) (*auth.OTPOk, *shared.ResponseOneMessage) {
	var err error
	reg := regexp.MustCompile(`^09\d{9}$`)
	if reg.MatchString(fields.Input) {
		_, err = a.userService.GetUserByMobile(fields.Input)
	} else {
		_, err = a.userService.GetUserByEmail(fields.Input)
	}
	if err != nil {
		if errors.Is(err, shared.ErrUsertNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusUnauthorized,
				ErrMessage: shared.ErrInvalidMobile,
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
	generatedCode, ok, result := a.otpService.StoreCode(otp.FieldOTPStore{
		Prefix: "reset:",
		Key:    fields.Input,
	})
	if result != nil {
		return nil, result
	}
	userText := fmt.Sprintf(shared.TemplateSendVerifyOTP, generatedCode)
	if reg.MatchString(fields.Input) {
		result = a.smsService.SendToClient(sms.FieldSendClient{
			PhoneNumber: fields.Input,
			Text:        userText,
		})
		if result != nil {
			return nil, result
		}
	} else {
		a.emailService.SendToClient(email.FieldSendClient{
			Email: fields.Input,
			Text:  userText,
		})
	}
	return ok, nil
}
