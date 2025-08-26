package auth

type FieldSendOTP struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}
