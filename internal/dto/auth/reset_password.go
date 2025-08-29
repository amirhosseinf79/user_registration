package auth

type FieldSendResetPwd struct {
	Username string `json:"username" validate:"required"`
}

type FieldResetByOTP struct {
	FieldSendResetPwd
	Code        string `json:"code" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
