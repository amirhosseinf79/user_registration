package auth

type FieldUserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	UseOTP   bool   `json:"useOTP"`
}
