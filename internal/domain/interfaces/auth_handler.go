package interfaces

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	VerifyOTP(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}
