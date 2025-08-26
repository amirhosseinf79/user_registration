package service

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) RegisterUserByNumber(phoneNumber string) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	exists, err := u.userRepo.CheckMobileExists(phoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	userM := &model.User{
		PhoneNumber: phoneNumber,
	}
	if !exists {
		err = u.userRepo.Create(userM)
		if err != nil {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
	} else {
		userM, err = u.userRepo.GetByMobile(phoneNumber)
		if err != nil {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
	}
	userDetails := user.ResponseDetails{
		ID:           userM.ID,
		PhoneNumber:  userM.PhoneNumber,
		FirstName:    userM.FirstName,
		LastName:     userM.LastName,
		Email:        userM.Email,
		RegisteredAt: userM.CreatedAt,
	}
	return &userDetails, nil
}

func (u *userService) GetUserList(filter user.FilterUser) (*shared.ResponseList[user.ResponseDetails], *shared.ResponseOneMessage) {
	users, total, err := u.userRepo.GetAllByFilter(filter)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userList := []user.ResponseDetails{}
	for _, userM := range users {
		userDetails := user.ResponseDetails{
			ID:           userM.ID,
			PhoneNumber:  userM.PhoneNumber,
			FirstName:    userM.FirstName,
			LastName:     userM.LastName,
			Email:        userM.Email,
			RegisteredAt: userM.CreatedAt,
		}
		userList = append(userList, userDetails)
	}
	response := shared.NewResponseList(userList, int(total), filter.Page, filter.PageSize)
	return &response, nil
}

func (u *userService) GetUserDetailsByID(userID uint) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	userM, err := u.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared.ErrUsertNotFound,
				RealError:  err,
			})
			return nil, result
		}
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userDetails := user.ResponseDetails{
		ID:           userM.ID,
		PhoneNumber:  userM.PhoneNumber,
		FirstName:    userM.FirstName,
		LastName:     userM.LastName,
		Email:        userM.Email,
		RegisteredAt: userM.CreatedAt,
	}
	return &userDetails, nil
}

func (u *userService) UpdateUserProfile(userID uint, fields user.FieldUpdateDetails) (*user.ResponseDetails, *shared.ResponseOneMessage) {
	userM, err := u.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared.ErrUsertNotFound,
				RealError:  err,
			})
			return nil, result
		}
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
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
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userDetails := user.ResponseDetails{
		ID:           userM.ID,
		PhoneNumber:  userM.PhoneNumber,
		FirstName:    userM.FirstName,
		LastName:     userM.LastName,
		Email:        userM.Email,
		RegisteredAt: userM.CreatedAt,
	}
	return &userDetails, nil
}
