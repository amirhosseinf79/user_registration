package user

import "github.com/gofiber/fiber/v2"

// SendUserVerifyMobile godoc
// @Summary Send Verify Code
// @Description Send Verify Code
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /api/v1/profile/send/verify-mobile-otp [post]
func (uh *userHandler) SendUserVerifyMobile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	response, err := uh.userService.SendVerifyMobile(userID)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
