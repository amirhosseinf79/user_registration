package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

func (a *authService) AutoLogin(field auth.FieldUserLogin) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	if field.UseOTP {
		return a.LoginByOTP(auth.FieldVerifyOTP{
			FieldSendOTP: auth.FieldSendOTP{
				PhoneNumber: field.Username,
			},
			Code: field.Password,
		})
	} else {
		if a.MatchMobile(field.Username) {
			return a.LoginByMobile(auth.FieldMobileLogin{
				PhoneNumber: field.Username,
				Password:    field.Password,
			})
		} else {
			return a.LoginByEmail(auth.FieldEmailLogin{
				Email:    field.Username,
				Password: field.Password,
			})
		}
	}
}
