package field_validator

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateEmail(ctx *fiber.Ctx) error {
	re := regexp.MustCompile(`^(.{3,})@(.{3,})\.(.{2,})$`)
	var fields user.FieldEmail
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if fields.Email != "" && !re.MatchString(fields.Email) {
		response := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidEmail,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	return ctx.Next()
}
