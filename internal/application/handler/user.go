package handler

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) interfaces.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// GetUserByID
// @Summary Say Hello
// @Description Returns Hello World message
// @Tags example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func (uh *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	userID, err := ctx.ParamsInt("userID", 0)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(err, fiber.StatusNotFound)
		return ctx.Status(statusCode).JSON(response)
	}
	userDetails, err := uh.userService.GetUserDetailsByID(uint(userID))
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(err, fiber.StatusNotFound)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(userDetails)
}

func (uh *userHandler) GetUsersList(ctx *fiber.Ctx) error {
	var filter dto.FilterUser
	ctx.QueryParser(&filter)
	userDetails, err := uh.userService.GetUserList(filter)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(err, fiber.StatusInternalServerError)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(userDetails)
}

func (uh *userHandler) UpdateProfileInfo(ctx *fiber.Ctx) error {
	var fields dto.UpdateUserDetails
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserProfile(userID, fields)
	if err != nil {
		statusCode, response := dto.NewDefaultRespose(err, fiber.StatusInternalServerError)
		return ctx.Status(statusCode).JSON(response)
	}
	return ctx.JSON(userDetails)
}
