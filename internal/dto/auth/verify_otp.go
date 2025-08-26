package auth

type FieldVerifyOTP struct {
	FieldSendOTP
	Code string `json:"code" validate:"required"`
}
