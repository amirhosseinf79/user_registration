package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	AutoLogin(ctx *fiber.Ctx) error
	RegisterByEmail(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	SendResetPassOTP(ctx *fiber.Ctx) error
	ResetPassWithOTP(ctx *fiber.Ctx) error
}

type AuthService interface {
	SendOTP(fields auth.FieldSendOTP) (*auth.OTPOk, *shared.ResponseOneMessage)
	RegisterByEmail(fields auth.FieldEmailRegister) (*auth.ResponseJWT, *shared.ResponseOneMessage)

	AutoLogin(field auth.FieldUserLogin) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	RefreshToken(refresh string) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	SendResetPasswerd(fields auth.FieldSendResetPwd) (*auth.OTPOk, *shared.ResponseOneMessage)
	ResetPassWithOTP(fields auth.FieldResetByOTP) *shared.ResponseOneMessage
}
