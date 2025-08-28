package interfaces

import "github.com/gofiber/fiber/v2"

type FieldValidatorMiddleware interface {
	ValidateMobile(ctx *fiber.Ctx) error
	ValidateEmailLogin(ctx *fiber.Ctx) error
	ValidateMobileLogin(ctx *fiber.Ctx) error
	ValidateVerifyCode(ctx *fiber.Ctx) error
	ValidateVerifyField(ctx *fiber.Ctx) error
	ValidateRefreshToken(ctx *fiber.Ctx) error
	ValidateEmailBody(ctx *fiber.Ctx) error
	ValidateNewPassword(ctx *fiber.Ctx) error
	ValidateRegister(ctx *fiber.Ctx) error
}

type AuthMiddleware interface {
	CheckToken(ctx *fiber.Ctx) error
}
