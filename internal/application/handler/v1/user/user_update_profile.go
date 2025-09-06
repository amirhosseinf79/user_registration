package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

// UpdateProfileInfo godoc
// @Summary Update user Profile
// @Description Update User Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body user.FieldUpdateDetails true "user"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/profile/update/info [patch]
func (uh *userHandler) UpdateProfileInfo(ctx *fiber.Ctx) error {
	var fields user.FieldUpdateDetails
	_ = ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserProfile(userID, fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
