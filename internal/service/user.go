package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
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

func (u *userService) GetUserList(filter dto.FilterUser) (*dto.ResponseList[dto.ResponseUserDetails], error) {
	users, total, err := u.userRepo.GetAllByFilter(filter)
	if err != nil {
		return nil, err
	}
	userList := []dto.ResponseUserDetails{}
	for _, user := range users {
		userDetails := dto.ResponseUserDetails{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			FieldEmail:  dto.FieldEmail{Email: user.Email},
		}
		userList = append(userList, userDetails)
	}
	response := dto.NewResponseList(userList, int(total), filter.Page, filter.PageSize)
	return &response, nil
}

func (u *userService) GetUserDetailsByID(userID uint) (*dto.ResponseUserDetails, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	userDetails := dto.ResponseUserDetails{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FieldEmail:  dto.FieldEmail{Email: user.Email},
	}
	return &userDetails, nil
}

func (u *userService) UpdateUserProfile(userID uint, fields dto.UpdateUserDetails) (*dto.ResponseUserDetails, error) {
	userM, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if fields.FirstName != "" {
		userM.FirstName = fields.FirstName
	}
	if fields.LastName != "" {
		userM.LastName = fields.LastName
	}
	if fields.Email != "" {
		userM.Email = fields.Email
	}

	err = u.userRepo.Update(userM)
	if err != nil {
		return nil, err
	}

	userDetails := dto.ResponseUserDetails{
		ID:          userM.ID,
		PhoneNumber: userM.PhoneNumber,
		FirstName:   userM.FirstName,
		LastName:    userM.LastName,
		FieldEmail:  dto.FieldEmail{Email: userM.Email},
	}
	return &userDetails, nil
}
