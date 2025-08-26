package user_response

import "time"

type Details struct {
	ID           uint      `json:"id"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email" query:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	RegisteredAt time.Time `json:"registeredAt"`
}
