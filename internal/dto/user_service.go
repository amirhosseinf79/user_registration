package dto

import "time"

type FieldEmail struct {
	Email string `json:"email" query:"email"`
}

type FilterUser struct {
	PhoneNumber string `query:"phoneNumber"`
	Name        string `query:"name"`
	FieldEmail
	FieldPagination
}

type ResponseUserDetails struct {
	ID           uint      `json:"id"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email" query:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	RegisteredAt time.Time `json:"registeredAt"`
}

type UpdateUserDetails struct {
	Email     string `json:"email" query:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ResponseUserList struct {
	Items []ResponseUserDetails `json:"items"`
	Meta  MetaPagination        `json:"meta"`
}
