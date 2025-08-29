package user

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type userHandler struct {
	userService interfaces.UserService1
}

func NewUserHandler(userService interfaces.UserService1) interfaces.UserHandler1 {
	return &userHandler{
		userService: userService,
	}
}
