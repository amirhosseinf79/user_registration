package auth_request

type FieldVerifyOTP struct {
	FieldSendOTP
	Code string `json:"code" validate:"required"`
}
