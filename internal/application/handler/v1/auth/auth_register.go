package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

// RegisterByEmail
// @Summary Register By Email
// @Description Register By Email
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldEmailRegister true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/auth/register [post]
func (ah *authHandler) RegisterByEmail(ctx *fiber.Ctx) error {
	var fields auth.FieldEmailRegister
	_ = ctx.BodyParser(&fields)
	response, err := ah.authService.RegisterByEmail(fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
