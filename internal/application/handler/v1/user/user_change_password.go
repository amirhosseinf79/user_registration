package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update User Password
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fields body user.FieldUpdatePassword true "Fields"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/profile/update/password [put]
func (uh *userHandler) UpdateUserPassword(ctx *fiber.Ctx) error {
	var fields user.FieldUpdatePassword
	_ = ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserPassword(userID, fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
