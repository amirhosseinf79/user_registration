package handler

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) interfaces.AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (ah *authHandler) SendOTP(ctx *fiber.Ctx) error {
	var fields dto.FieldAuthSendOTP
	ctx.BodyParser(&fields)
	err := ah.authService.SendOTP(fields)
	if errors.Is(err, dto.ErrSmsRateLimited) {
		statusCode, response := dto.NewDefaultRespose(err, fiber.StatusForbidden)
		return ctx.Status(statusCode).JSON(response)
	}
	statusCode, response := dto.NewDefaultRespose(err, fiber.StatusBadRequest)
	return ctx.Status(statusCode).JSON(response)
}

func (ah *authHandler) VerifyOTP(ctx *fiber.Ctx) error {
	var fields dto.FieldAuthVerifyOTP
	ctx.BodyParser(&fields)
	response, err := ah.authService.VerifyOTP(fields)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(dto.ErrUnauthorized, fiber.StatusUnauthorized)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(response)
}

func (ah *authHandler) RefreshToken(ctx *fiber.Ctx) error {
	var fields dto.FieldRefreshToken
	ctx.BodyParser(&fields)
	response, err := ah.authService.RefreshToken(fields.RefreshToken)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(dto.ErrInvalidToken, fiber.StatusUnauthorized)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(response)
}
