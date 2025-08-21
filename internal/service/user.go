package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) RegisterUserByNumber(phoneNumber string) (*model.User, error) {
	exists, err := u.userRepo.CheckMobileExists(phoneNumber)
	if err != nil {
		return nil, err
	}

	if !exists {
		user := &model.User{
			PhoneNumber: phoneNumber,
		}
		err = u.userRepo.Create(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	user, err := u.userRepo.GetByMobile(phoneNumber)
	if err != nil {
		return nil, err
	}
	return user, nil
}
