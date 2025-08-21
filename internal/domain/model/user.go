package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	PhoneNumber string
	FirstName   string
	LastName    string
	Email       string
}
