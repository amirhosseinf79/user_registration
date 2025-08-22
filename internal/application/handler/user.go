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

// @Summary Get user details
// @Description Get User Details by ID
// @Tags user
// @Accept json
// @Produce json
// @Param userID path int true "UserID"
// @Success 200 {object} dto.ResponseUserDetails
// @Failure 500 {object} dto.responseOneMessage
// @Router /user/{userID} [get]
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

// @Summary Get all users
// @Description Get list of all users
// @Tags user
// @Accept json
// @Produce json
// @Param filters query dto.FilterUser false "Filters"
// @Success 200 {object} dto.ResponseUserList
// @Failure 500 {object} dto.responseOneMessage
// @Router /user/all [get]
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

// @Summary Update user Profile
// @Description Update User Profile
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.UpdateUserDetails true "user"
// @Success 200 {object} dto.ResponseUserDetails
// @Failure 500 {object} dto.responseOneMessage
// @Router /profile/update [post]
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
