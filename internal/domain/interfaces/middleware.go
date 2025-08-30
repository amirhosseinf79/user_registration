package interfaces

import "github.com/gofiber/fiber/v2"

type FieldValidatorMiddleware1 interface {
	ValidateSendResetPassOTP(ctx *fiber.Ctx) error
	ValidateNewPassword(ctx *fiber.Ctx) error
	ValidateMobile(ctx *fiber.Ctx) error
	ValidateEmail(ctx *fiber.Ctx) error
	ValidateVerifyCode(ctx *fiber.Ctx) error
	ValidateRefreshToken(ctx *fiber.Ctx) error
	ValidateRegister(ctx *fiber.Ctx) error
	ValidateLogin(ctx *fiber.Ctx) error
	ValidatePassword(ctx *fiber.Ctx) error
	ValidateUpdateMobile(ctx *fiber.Ctx) error
}

type AuthMiddleware1 interface {
	CheckToken(ctx *fiber.Ctx) error
}
