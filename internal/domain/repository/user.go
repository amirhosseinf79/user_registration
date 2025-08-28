package repository

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
)

type UserRepository interface {
	GetAllByFilter(filter user.FilterUser) ([]*model.User, int64, error)
	GetByEmail(email string) (*model.User, error)
	GetByMobile(mobile string) (*model.User, error)
	CheckMobileExists(mobile string) (bool, error)
	CheckEmailExists(email string) (bool, error)
	GetByID(userID uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userID uint) error
}
