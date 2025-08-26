package handler

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	user_request "github.com/amirhosseinf79/user_registration/internal/dto/user/request"
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
// @Success 200 {object} user_response.Details
// @Failure 404 {object} shared_dto.ResponseOneMessage
// @Failure 500 {object} shared_dto.ResponseOneMessage
// @Router /user/{userID} [get]
func (uh *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	userID, _ := ctx.ParamsInt("userID", 0)
	userDetails, err := uh.userService.GetUserDetailsByID(uint(userID))
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Get all users
// @Description Get list of all users
// @Tags user
// @Accept json
// @Produce json
// @Param filters query user_request.FilterUser false "Filters"
// @Success 200 {object} user_response.List
// @Failure 500 {object} shared_dto.ResponseOneMessage
// @Router /user/all [get]
func (uh *userHandler) GetUsersList(ctx *fiber.Ctx) error {
	var filter user_request.FilterUser
	ctx.QueryParser(&filter)
	userDetails, err := uh.userService.GetUserList(filter)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Update user Profile
// @Description Update User Profile
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body user_request.UpdateDetails true "user"
// @Success 200 {object} user_response.Details
// @Failure 500 {object} shared_dto.ResponseOneMessage
// @Router /profile/update [patch]
func (uh *userHandler) UpdateProfileInfo(ctx *fiber.Ctx) error {
	var fields user_request.UpdateDetails
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserProfile(userID, fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
