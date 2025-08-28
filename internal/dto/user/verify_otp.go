package user

type FieldVerifyOTP struct {
	Code string `json:"code" validate:"required"`
}
