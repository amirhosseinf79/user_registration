package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserByID(ctx *fiber.Ctx) error
	GetUsersList(ctx *fiber.Ctx) error
	GetUserProfile(ctx *fiber.Ctx) error
	UpdateProfileInfo(ctx *fiber.Ctx) error
	UpdateUserPassword(ctx *fiber.Ctx) error

	SendUserVerifyMobile(ctx *fiber.Ctx) error
	SendUserVerifyEmail(ctx *fiber.Ctx) error
	VerifyUserMobile(ctx *fiber.Ctx) error
	VerifyUserEmail(ctx *fiber.Ctx) error
}

type UserService interface {
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByMobile(mobile string) (*model.User, error)

	GetUserList(filter user.FilterUser) (*shared.ResponseList[user.ResponseDetails], *shared.ResponseOneMessage)
	GetUserDetailsByMobile(mobile string) (*user.ResponseDetails, *shared.ResponseOneMessage)
	GetUserDetailsByEmail(email string) (*user.ResponseDetails, *shared.ResponseOneMessage)
	GetUserDetailsByID(userID uint) (*user.ResponseDetails, *shared.ResponseOneMessage)

	CheckUserEmailExists(email string) *shared.ResponseOneMessage
	CheckUserMobileExists(mobile string) *shared.ResponseOneMessage

	RegisterUserByNumber(phoneNumber string) (*user.ResponseDetails, *shared.ResponseOneMessage)
	RegisterUserByEmail(fields auth.FieldEmailRegister) (*user.ResponseDetails, *shared.ResponseOneMessage)

	UpdateUserProfile(userID uint, fields user.FieldUpdateDetails) (*user.ResponseDetails, *shared.ResponseOneMessage)
	UpdateUserPassword(userID uint, fields user.FieldUpdatePassword) (*user.ResponseDetails, *shared.ResponseOneMessage)
	ResetUserPasswordByInfo(userInfo, newPassword string) (*user.ResponseDetails, *shared.ResponseOneMessage)

	SendVerifyMobile(userID uint) (*auth.OTPOk, *shared.ResponseOneMessage)
	SendVerifyEmail(userID uint) (*auth.OTPOk, *shared.ResponseOneMessage)
	VerifyUserMobile(userID uint, code string) (*user.ResponseDetails, *shared.ResponseOneMessage)
	VerifyUserEmail(userID uint, code string) (*user.ResponseDetails, *shared.ResponseOneMessage)
}
