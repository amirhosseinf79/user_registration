package model

import (
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber string
	FirstName   string
	LastName    string
	Email       string
	Password    string
}

func (u *User) ValidatePassword(password string) bool {
	return pkg.ComparePassword(password, u.Password)
}
