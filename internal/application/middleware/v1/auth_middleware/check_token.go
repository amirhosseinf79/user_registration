package auth_middleware

import (
	"strings"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (am *authMiddleware) CheckToken(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization", "")
	if token == "" || !strings.Contains(strings.ToLower(token), strings.ToLower(am.prefix)) {
		response := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidToken,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	userID, err := am.jwtService.GetUserIDByAccessToken(token[len(am.prefix):])
	if err != nil {
		response := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrUnauthorized,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	ctx.Locals("userID", userID)
	return ctx.Next()
}
