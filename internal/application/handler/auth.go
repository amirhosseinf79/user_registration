package handler

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	auth_request "github.com/amirhosseinf79/user_registration/internal/dto/auth/request"
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
// @Param fields body auth_request.FieldSendOTP true "Fields"
// @Success 200 {array} shared_dto.ResponseOneMessage
// @Failure 400 {object} shared_dto.ResponseOneMessage
// @Failure 403 {object} shared_dto.ResponseOneMessage
// @Router /auth/send-otp [post]
func (ah *authHandler) SendOTP(ctx *fiber.Ctx) error {
	var fields auth_request.FieldSendOTP
	ctx.BodyParser(&fields)
	response := ah.authService.SendOTP(fields)
	return ctx.Status(response.Code).JSON(response)
}

// @Summary Verify OTP
// @Description Verify OTP & Login or Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param fields body auth_request.FieldVerifyOTP true "Fields"
// @Success 200 {array} auth_response.JWT
// @Failure 401 {object} shared_dto.ResponseOneMessage
// @Router /auth/verify-otp [post]
func (ah *authHandler) LoginByOTP(ctx *fiber.Ctx) error {
	var fields auth_request.FieldVerifyOTP
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
// @Param fields body auth_request.FieldRefreshToken true "Fields"
// @Success 200 {array} auth_response.JWT
// @Failure 401 {object} shared_dto.ResponseOneMessage
// @Router /auth/refresh-token [post]
func (ah *authHandler) RefreshToken(ctx *fiber.Ctx) error {
	var fields auth_request.FieldRefreshToken
	ctx.BodyParser(&fields)
	response, err := ah.authService.RefreshToken(fields.RefreshToken)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
