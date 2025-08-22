package interfaces

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	SendOTP(ctx *fiber.Ctx) error
	LoginByOTP(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}
