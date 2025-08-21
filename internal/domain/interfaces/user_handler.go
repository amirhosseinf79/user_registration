package interfaces

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	GetUserByID(ctx *fiber.Ctx) error
	GetUsersList(ctx *fiber.Ctx) error
	UpdateProfileInfo(ctx *fiber.Ctx) error
}
