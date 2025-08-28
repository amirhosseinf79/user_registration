package handler

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
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
// @Param fields body auth.FieldSendOTP true "Fields"
// @Success 200 {object} auth.OTPOkMock
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 403 {object} shared.ResponseOneMessage
// @Router /auth/send-otp [post]
func (ah *authHandler) SendOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldSendOTP
	ctx.BodyParser(&fields)
	response, err := ah.authService.SendOTP(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.Status(response.Code).JSON(response)
}

// @Summary Verify OTP
// @Description Verify OTP & Login or Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldVerifyOTP true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/verify-otp [post]
func (ah *authHandler) LoginByOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldVerifyOTP
	ctx.BodyParser(&fields)
	response, err := ah.authService.LoginByOTP(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldRefreshToken true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/refresh-token [post]
func (ah *authHandler) RefreshToken(ctx *fiber.Ctx) error {
	var fields auth.FieldRefreshToken
	ctx.BodyParser(&fields)
	response, err := ah.authService.RefreshToken(fields.RefreshToken)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
