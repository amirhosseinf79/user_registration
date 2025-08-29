package middleware

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateEmailLogin(ctx *fiber.Ctx) error {
	var fields auth.FieldEmailLogin
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}
