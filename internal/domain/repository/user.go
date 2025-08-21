package repository

import "github.com/amirhosseinf79/user_registration/internal/domain/model"

type UserRepository interface {
	GetByMobile(mobile string) (*model.User, error)
	CheckMobileExists(mobile string) (bool, error)
	GetByID(userID uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userID uint) error
}
