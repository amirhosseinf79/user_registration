package auth_request

type FieldSendOTP struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}
