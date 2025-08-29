package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

// @Summary Verify Mobile
// @Description verify Mobile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fields body user.FieldVerifyOTP true "Fields"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/profile/verify/email [post]
func (ah *userHandler) VerifyUserEmail(ctx *fiber.Ctx) error {
	var fields user.FieldVerifyOTP
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	response, err := ah.userService.VerifyUserEmail(userID, fields.Code)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
