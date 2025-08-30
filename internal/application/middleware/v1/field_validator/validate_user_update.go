package field_validator

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateUpdateMobile(ctx *fiber.Ctx) error {
	var fields user.FieldUpdateDetails
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if fields.PhoneNumber != "" {
		return fv.ValidateMobile(ctx)
	}
	return ctx.Next()
}
