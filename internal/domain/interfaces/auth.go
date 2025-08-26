package interfaces

import (
	auth_request "github.com/amirhosseinf79/user_registration/internal/dto/auth/request"
	auth_response "github.com/amirhosseinf79/user_registration/internal/dto/auth/response"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	LoginByOTP(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}

type AuthService interface {
	SendOTP(fields auth_request.FieldSendOTP) *shared_dto.ResponseOneMessage
	LoginByOTP(fields auth_request.FieldVerifyOTP) (*auth_response.JWT, *shared_dto.ResponseOneMessage)
	RefreshToken(refresh string) (*auth_response.JWT, *shared_dto.ResponseOneMessage)
}
