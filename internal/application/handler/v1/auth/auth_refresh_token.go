package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/gofiber/fiber/v2"
)

// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param fields body auth.FieldRefreshToken true "Fields"
// @Success 200 {array} auth.ResponseJWT
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/auth/refresh-token [post]
func (ah *authHandler) RefreshToken(ctx *fiber.Ctx) error {
	var fields auth.FieldRefreshToken
	ctx.BodyParser(&fields)
	response, err := ah.authService.RefreshToken(fields.RefreshToken)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
