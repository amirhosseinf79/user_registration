package interfaces

import "github.com/gofiber/fiber/v3"

type FieldValidatorMiddleware interface {
	ValidateMobile(ctx fiber.Ctx) error
	ValidateCode(ctx fiber.Ctx) error
	ValidateRefreshToken(ctx fiber.Ctx) error
}
