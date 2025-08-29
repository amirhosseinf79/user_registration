package field_validator

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateRefreshToken(ctx *fiber.Ctx) error {
	var fields auth.FieldRefreshToken
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}
