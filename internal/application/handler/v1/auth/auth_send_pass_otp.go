package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

// SendResetPassOTP
// @Summary Send Reset Code
// @Description Send Reset Password Code
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldSendResetPwd true "Fields"
// @Success 200 {array} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/auth/password/send-code [post]
func (ah *authHandler) SendResetPassOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldSendResetPwd
	_ = ctx.BodyParser(&fields)
	response, err := ah.authService.SendResetPassword(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
