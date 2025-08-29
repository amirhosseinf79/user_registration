package otp

type FieldVerifyOTP struct {
	FieldOTPStore
	Code string `json:"code" validate:"required"`
}
