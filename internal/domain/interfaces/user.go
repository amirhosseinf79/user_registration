package interfaces

import (
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	user_request "github.com/amirhosseinf79/user_registration/internal/dto/user/request"
	user_response "github.com/amirhosseinf79/user_registration/internal/dto/user/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserByID(ctx *fiber.Ctx) error
	GetUsersList(ctx *fiber.Ctx) error
	UpdateProfileInfo(ctx *fiber.Ctx) error
}

type UserService interface {
	RegisterUserByNumber(phoneNumber string) (*user_response.Details, *shared_dto.ResponseOneMessage)
	GetUserList(filter user_request.FilterUser) (*shared_dto.ResponseList[user_response.Details], *shared_dto.ResponseOneMessage)
	UpdateUserProfile(userID uint, fields user_request.UpdateDetails) (*user_response.Details, *shared_dto.ResponseOneMessage)
	GetUserDetailsByID(userID uint) (*user_response.Details, *shared_dto.ResponseOneMessage)
}
