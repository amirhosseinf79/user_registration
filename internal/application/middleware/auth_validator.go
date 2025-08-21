package middleware

import (
	"strings"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type authMiddleware struct {
	jwtService interfaces.JWTService
	prefix     string
}

func NewAuthMiddleware(jwtService interfaces.JWTService) interfaces.AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
		prefix:     "Bearer ",
	}
}

func (am *authMiddleware) CheckToken(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization", "")
	if token == "" || !strings.Contains(strings.ToLower(token), strings.ToLower(am.prefix)) {
		statuscode, response := dto.NewDefaultRespose(dto.ErrInvalidToken, fiber.StatusUnauthorized)
		return ctx.Status(statuscode).JSON(response)
	}
	userID, err := am.jwtService.GetUserIDByAccessToken(token[len(am.prefix):])
	if err != nil {
		statuscode, response := dto.NewDefaultRespose(dto.ErrUnauthorized, fiber.StatusUnauthorized)
		return ctx.Status(statuscode).JSON(response)
	}
	ctx.Locals("userID", userID)
	return ctx.Next()
}
