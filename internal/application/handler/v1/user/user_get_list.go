package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

// GetUsersList godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags User
// @Accept json
// @Produce json
// @Param filters query user.FilterUser false "Filters"
// @Success 200 {object} user.ResponseList
// @Router /api/v1/user/all [get]
func (uh *userHandler) GetUsersList(ctx *fiber.Ctx) error {
	var filter user.FilterUser
	_ = ctx.QueryParser(&filter)
	userDetails, err := uh.userService.GetUserList(filter)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
