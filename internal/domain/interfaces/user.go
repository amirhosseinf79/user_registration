package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserByID(ctx *fiber.Ctx) error
	GetUsersList(ctx *fiber.Ctx) error
	UpdateProfileInfo(ctx *fiber.Ctx) error
}

type UserService interface {
	RegisterUserByNumber(phoneNumber string) (*user.ResponseDetails, *shared.ResponseOneMessage)
	GetUserList(filter user.FilterUser) (*shared.ResponseList[user.ResponseDetails], *shared.ResponseOneMessage)
	UpdateUserProfile(userID uint, fields user.FieldUpdateDetails) (*user.ResponseDetails, *shared.ResponseOneMessage)
	GetUserDetailsByID(userID uint) (*user.ResponseDetails, *shared.ResponseOneMessage)
}
