package middleware

import (
	"strings"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
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
		response := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared_dto.ErrInvalidToken,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	userID, err := am.jwtService.GetUserIDByAccessToken(token[len(am.prefix):])
	if err != nil {
		response := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared_dto.ErrUnauthorized,
		})
		return ctx.Status(response.Code).JSON(response)
	}
	ctx.Locals("userID", userID)
	return ctx.Next()
}
