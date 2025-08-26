package service

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	user_request "github.com/amirhosseinf79/user_registration/internal/dto/user/request"
	user_response "github.com/amirhosseinf79/user_registration/internal/dto/user/response"
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

func (u *userService) RegisterUserByNumber(phoneNumber string) (*user_response.Details, *shared_dto.ResponseOneMessage) {
	exists, err := u.userRepo.CheckMobileExists(phoneNumber)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}

	user := &model.User{
		PhoneNumber: phoneNumber,
	}
	if !exists {
		err = u.userRepo.Create(user)
		if err != nil {
			result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared_dto.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
	} else {
		user, err = u.userRepo.GetByMobile(phoneNumber)
		if err != nil {
			result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared_dto.ErrInternalServerError,
				RealError:  err,
			})
			return nil, result
		}
	}
	userDetails := user_response.Details{
		ID:           user.ID,
		PhoneNumber:  user.PhoneNumber,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		RegisteredAt: user.CreatedAt,
	}
	return &userDetails, nil
}

func (u *userService) GetUserList(filter user_request.FilterUser) (*shared_dto.ResponseList[user_response.Details], *shared_dto.ResponseOneMessage) {
	users, total, err := u.userRepo.GetAllByFilter(filter)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userList := []user_response.Details{}
	for _, user := range users {
		userDetails := user_response.Details{
			ID:           user.ID,
			PhoneNumber:  user.PhoneNumber,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			RegisteredAt: user.CreatedAt,
		}
		userList = append(userList, userDetails)
	}
	response := shared_dto.NewResponseList(userList, int(total), filter.Page, filter.PageSize)
	return &response, nil
}

func (u *userService) GetUserDetailsByID(userID uint) (*user_response.Details, *shared_dto.ResponseOneMessage) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared_dto.ErrUsertNotFound,
				RealError:  err,
			})
			return nil, result
		}
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userDetails := user_response.Details{
		ID:           user.ID,
		PhoneNumber:  user.PhoneNumber,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		RegisteredAt: user.CreatedAt,
	}
	return &userDetails, nil
}

func (u *userService) UpdateUserProfile(userID uint, fields user_request.UpdateDetails) (*user_response.Details, *shared_dto.ResponseOneMessage) {
	userM, err := u.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
				ErrStatus:  fiber.StatusNotFound,
				ErrMessage: shared_dto.ErrUsertNotFound,
				RealError:  err,
			})
			return nil, result
		}
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
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
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	userDetails := user_response.Details{
		ID:           userM.ID,
		PhoneNumber:  userM.PhoneNumber,
		FirstName:    userM.FirstName,
		LastName:     userM.LastName,
		Email:        userM.Email,
		RegisteredAt: userM.CreatedAt,
	}
	return &userDetails, nil
}
