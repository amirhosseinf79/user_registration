package user

import "github.com/gofiber/fiber/v2"

// @Summary Get user details
// @Description Get User Details by ID
// @Tags User
// @Accept json
// @Produce json
// @Param userID path int true "UserID"
// @Success 200 {object} user.ResponseDetails
// @Failure 404 {object} shared.ResponseOneMessage
// @Router /api/v1/user/{userID} [get]
func (uh *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	userID, _ := ctx.ParamsInt("userID", 0)
	userDetails, err := uh.userService.GetUserDetailsByID(uint(userID))
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}
