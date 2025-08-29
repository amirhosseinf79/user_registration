package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

// @Summary Reset Password
// @Description change passwrod by sent OTP. Username could be Number or Email
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldResetByOTP true "Fields"
// @Success 200 {array} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/auth/password/reset [put]
func (ah *authHandler) ResetPassWithOTP(ctx *fiber.Ctx) error {
	var fields auth.FieldResetByOTP
	ctx.BodyParser(&fields)
	response := ah.authService.ResetPassWithOTP(fields)
	return ctx.Status(response.Code).JSON(response)
}
