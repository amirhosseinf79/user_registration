package user

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type ResponseDetails struct {
	ID             uint      `json:"id"`
	PhoneNumber    string    `json:"phoneNumber"`
	Email          string    `json:"email" query:"email"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	RegisteredAt   time.Time `json:"registeredAt"`
	HasPassword    bool      `json:"hasPassword"`
	MobileVerified bool      `json:"mobileVerified"`
	EmailVerified  bool      `json:"emailVerified"`
}

func NewUserResponse(userM *model.User) *ResponseDetails {
	return &ResponseDetails{
		ID:             userM.ID,
		PhoneNumber:    userM.PhoneNumber,
		FirstName:      userM.FirstName,
		LastName:       userM.LastName,
		Email:          userM.Email,
		RegisteredAt:   userM.CreatedAt,
		HasPassword:    userM.Password != "",
		MobileVerified: userM.MobileVerified,
		EmailVerified:  userM.EmailVerified,
	}
}
