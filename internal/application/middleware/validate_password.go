package middleware

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateNewPassword(ctx *fiber.Ctx) error {
	re1 := regexp.MustCompile(`.{10}`)
	re2 := regexp.MustCompile(`[@#$%^&*]+`)
	var fields user.FieldUpdatePassword
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if !re1.MatchString(fields.NewPassword) || !re2.MatchString(fields.NewPassword) {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrPasswordValidation,
		})
		return ctx.Status(result.Code).JSON(result)
	}
	return ctx.Next()
}
