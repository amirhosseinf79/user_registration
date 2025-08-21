package interfaces

import "github.com/amirhosseinf79/user_registration/internal/domain/model"

type UserService interface {
	RegisterUserByNumber(phoneNumber string) (*model.User, error)
}
