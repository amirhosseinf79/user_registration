package middleware

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

type fieldsValidatorMiddleware struct{}

func NewFieldValidator() interfaces.FieldValidatorMiddleware {
	return &fieldsValidatorMiddleware{}
}

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

func (fv *fieldsValidatorMiddleware) ValidateLogin(ctx *fiber.Ctx) error {
	var fields auth.FieldPassLogin
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}

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

func (fv *fieldsValidatorMiddleware) ValidateCode(ctx *fiber.Ctx) error {
	var fields auth.FieldVerifyOTP
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}

func (fv *fieldsValidatorMiddleware) ValidateRefreshToken(ctx *fiber.Ctx) error {
	var fields auth.FieldRefreshToken
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	return ctx.Next()
}

func (fv *fieldsValidatorMiddleware) ValidateEmailBody(ctx *fiber.Ctx) error {
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
