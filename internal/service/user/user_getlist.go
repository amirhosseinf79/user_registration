package user

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

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
			HasPassword:  userM.Password != "",
		}
		userList = append(userList, userDetails)
	}
	response := shared.NewResponseList(userList, int(total), filter.Page, filter.PageSize)
	return &response, nil
}
