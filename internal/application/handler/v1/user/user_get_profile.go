package user

import "github.com/gofiber/fiber/v2"

// GetUserProfile godoc
// @Summary Get user profile
// @Description Get User profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} user.ResponseDetails
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/profile [get]
func (uh *userHandler) GetUserProfile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.GetUserDetailsByID(userID)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
