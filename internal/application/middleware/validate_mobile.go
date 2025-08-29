package middleware

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateMobile(ctx *fiber.Ctx) error {
	re := regexp.MustCompile(`^09\d{9}$`)
	var fields auth.FieldSendOTP
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if !re.MatchString(fields.PhoneNumber) {
		response := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidMobile,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	return ctx.Next()
}
