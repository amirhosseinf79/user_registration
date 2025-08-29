package interfaces

import "github.com/gofiber/fiber/v2"

type FieldValidatorMiddleware interface {
	ValidateSendResetPassOTP(ctx *fiber.Ctx) error
	ValidateNewPassword(ctx *fiber.Ctx) error
	ValidateMobile(ctx *fiber.Ctx) error
	ValidateEmail(ctx *fiber.Ctx) error
	ValidateVerifyCode(ctx *fiber.Ctx) error
	ValidateRefreshToken(ctx *fiber.Ctx) error
	ValidateRegister(ctx *fiber.Ctx) error
	ValidateLogin(ctx *fiber.Ctx) error
}

type AuthMiddleware interface {
	CheckToken(ctx *fiber.Ctx) error
}
