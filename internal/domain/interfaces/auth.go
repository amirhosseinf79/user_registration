package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	RegisterByEmail(ctx *fiber.Ctx) error
	LoginByOTP(ctx *fiber.Ctx) error
	LoginByMobile(ctx *fiber.Ctx) error
	LoginByEmail(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}

type AuthService interface {
	SendOTP(fields auth.FieldSendOTP) (*auth.OTPOk, *shared.ResponseOneMessage)
	RegisterByEmail(fields auth.FieldEmailRegister) (*auth.ResponseJWT, *shared.ResponseOneMessage)

	LoginByOTP(fields auth.FieldVerifyOTP) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	LoginByEmail(fields auth.FieldEmailLogin) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	LoginByMobile(fields auth.FieldMobileLogin) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	RefreshToken(refresh string) (*auth.ResponseJWT, *shared.ResponseOneMessage)
}
