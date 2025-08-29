package field_validator

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (fv *fieldsValidatorMiddleware) ValidateLogin(ctx *fiber.Ctx) error {
	var fields auth.FieldUserLogin
	reMobile := regexp.MustCompile(`^09\d{9}$`)
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if fields.UseOTP && !reMobile.MatchString(fields.Username) {
		response := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidMobile,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	return ctx.Next()
}
