package auth

type FieldSendResetPwd struct {
	Input string `json:"input" validate:"required"`
}

type FieldResetByOTP struct {
	FieldSendResetPwd
	Code        string `json:"code" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
