package middleware

import (
	"regexp"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

type fieldsValidatorMiddleware struct{}

func NewFieldValidator() interfaces.FieldValidatorMiddleware {
	return &fieldsValidatorMiddleware{}
}

func (fv *fieldsValidatorMiddleware) ValidateMobile(ctx *fiber.Ctx) error {
	re := regexp.MustCompile(`^09\d{9}$`)
	var fields dto.FieldAuthSendOTP
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		statusCode, _ := dto.NewDefaultRespose(err, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	if !re.MatchString(fields.PhoneNumber) {
		statusCode, response := dto.NewDefaultRespose(dto.ErrInvalidMobile, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.Next()
}

func (fv *fieldsValidatorMiddleware) ValidateCode(ctx *fiber.Ctx) error {
	var fields dto.FieldAuthVerifyOTP
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		statusCode, _ := dto.NewDefaultRespose(err, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.Next()
}

func (fv *fieldsValidatorMiddleware) ValidateRefreshToken(ctx *fiber.Ctx) error {
	var fields dto.FieldRefreshToken
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		statusCode, _ := dto.NewDefaultRespose(err, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.Next()
}

func (fv *fieldsValidatorMiddleware) ValidateEmailBody(ctx *fiber.Ctx) error {
	re := regexp.MustCompile(`^(.{3,})@(.{3,})\.(.{2,})$`)
	var fields dto.FieldEmail
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		statusCode, _ := dto.NewDefaultRespose(err, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	if fields.Email != "" && !re.MatchString(fields.Email) {
		statusCode, response := dto.NewDefaultRespose(dto.ErrInvalidEmail, fiber.StatusBadRequest)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.Next()
}
