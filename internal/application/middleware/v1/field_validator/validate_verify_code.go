package field_validator

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateVerifyCode(ctx *fiber.Ctx) error {
	var fields user.FieldVerifyOTP
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}
