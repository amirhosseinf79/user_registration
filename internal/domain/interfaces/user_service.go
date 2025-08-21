package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto"
)

type UserService interface {
	RegisterUserByNumber(phoneNumber string) (*model.User, error)
	GetUserList(filter dto.FilterUser) (*dto.ResponseList[dto.ResponseUserDetails], error)
	UpdateUserProfile(userID uint, fields dto.UpdateUserDetails) (*dto.ResponseUserDetails, error)
	GetUserDetailsByID(userID uint) (*dto.ResponseUserDetails, error)
}
