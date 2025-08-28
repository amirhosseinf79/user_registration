package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	LoginByOTP(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}

type AuthService interface {
	SendOTP(fields auth.FieldSendOTP) (*auth.OTPOk, *shared.ResponseOneMessage)
	LoginByOTP(fields auth.FieldVerifyOTP) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	RefreshToken(refresh string) (*auth.ResponseJWT, *shared.ResponseOneMessage)
}
