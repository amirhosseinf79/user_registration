package model

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/enum"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserType       enum.UserType ``
	Email          string        `gorm:"uniqueIndex:email"`
	PhoneNumber    string        `gorm:"uniqueIndex:mobile"`
	FirstName      string        ``
	LastName       string        ``
	Password       string        ``
	MobileVerified bool          ``
	EmailVerified  bool          ``
}

func (u *User) ValidatePassword(password string) bool {
	return pkg.ComparePassword(password, u.Password)
}
