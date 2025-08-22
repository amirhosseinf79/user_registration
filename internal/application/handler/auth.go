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

// @Summary Send OTP
// @Description Send OTP
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body dto.FieldAuthSendOTP true "Fields"
// @Success 200 {array} dto.responseOneMessage
// @Failure 400 {object} dto.responseOneMessage
// @Failure 403 {object} dto.responseOneMessage
// @Router /auth/send-otp [post]
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

// @Summary Verify OTP
// @Description Verify OTP & Login or Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body dto.FieldAuthVerifyOTP true "Fields"
// @Success 200 {array} dto.ResponseAuthOk
// @Failure 401 {object} dto.responseOneMessage
// @Router /auth/verify-otp [post]
func (ah *authHandler) LoginByOTP(ctx *fiber.Ctx) error {
	var fields dto.FieldAuthVerifyOTP
	ctx.BodyParser(&fields)
	response, err := ah.authService.LoginByOTP(fields)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(dto.ErrUnauthorized, fiber.StatusUnauthorized)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(response)
}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body dto.FieldRefreshToken true "Fields"
// @Success 200 {array} dto.ResponseAuthOk
// @Failure 401 {object} dto.responseOneMessage
// @Router /auth/refresh-token [post]
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
