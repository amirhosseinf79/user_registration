package interfaces

import "github.com/gofiber/fiber/v3"

type AuthMiddleware interface {
	CheckToken(ctx fiber.Ctx) error
}
