package user

import (
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
)

func (a *userService) SendVerifyOTP(userID uint) (*auth.OTPOk, *shared.ResponseOneMessage) {
	userM, err := a.GetUserDetailsByID(userID)
	if err != nil {
		return nil, err
	}
	generatedCode, response, err := a.otpService.StoreCode(auth.FieldSendOTP{
		PhoneNumber: fmt.Sprintf("V:%v", userM.PhoneNumber)},
	)
	if err != nil {
		return nil, err
	}
	smsText := fmt.Sprintf(shared.TemplateSendVerifyOTP, generatedCode)
	err = a.smsService.SendToClient(sms.FieldSendClient{
		PhoneNumber: userM.PhoneNumber,
		Text:        smsText,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
