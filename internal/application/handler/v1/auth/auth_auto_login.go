package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

// AutoLogin
// @Summary Login
// @Description Login By Email, Mobile or OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldUserLogin true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/auth/login [post]
func (ah *authHandler) AutoLogin(ctx *fiber.Ctx) error {
	var fields auth.FieldUserLogin
	_ = ctx.BodyParser(&fields)
	response, err := ah.authService.AutoLogin(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
