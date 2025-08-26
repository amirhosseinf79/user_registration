package user

import "time"

type ResponseDetails struct {
	ID           uint      `json:"id"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email" query:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	RegisteredAt time.Time `json:"registeredAt"`
}
