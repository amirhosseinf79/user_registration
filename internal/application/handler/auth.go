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
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldSendOTP true "Fields"
// @Success 200 {object} auth.OTPOkMock
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
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

// @Summary Login
// @Description Login By Email, Mobile or OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldUserLogin true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/login [post]
func (ah *authHandler) AutoLogin(ctx *fiber.Ctx) error {
	var fields auth.FieldUserLogin
	ctx.BodyParser(&fields)
	response, err := ah.authService.AutoLogin(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
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

// @Summary Register By Email
// @Description Register By Email
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldEmailRegister true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/register [post]
func (ah *authHandler) RegisterByEmail(ctx *fiber.Ctx) error {
	var fields auth.FieldEmailRegister
	ctx.BodyParser(&fields)
	response, err := ah.authService.RegisterByEmail(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}

// @Summary Send Reset Code
// @Description Send Reset Password Code
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldSendResetPwd true "Fields"
// @Success 200 {array} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/password/send-code [post]
func (ah *authHandler) SendResetPassOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldSendResetPwd
	ctx.BodyParser(&fields)
	response, err := ah.authService.SendResetPasswerd(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}

// @Summary Reset Password
// @Description change passwrod by sent OTP. Username could be Number or Email
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldResetByOTP true "Fields"
// @Success 200 {array} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /auth/password/reset [put]
func (ah *authHandler) ResetPassWithOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldResetByOTP
	ctx.BodyParser(&fields)
	response := ah.authService.ResetPassWithOTP(fields)
	return ctx.Status(response.Code).JSON(response)
}
