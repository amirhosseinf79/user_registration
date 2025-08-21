package interfaces

import "github.com/gofiber/fiber/v2"

type AuthMiddleware interface {
	CheckToken(ctx *fiber.Ctx) error
}
