package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

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
// @Router /api/v1/auth/send-otp [post]
func (ah *authHandler) SendOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldSendOTP
	ctx.BodyParser(&fields)
	response, err := ah.authService.SendOTP(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.Status(response.Code).JSON(response)
}
