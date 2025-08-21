package dto

type FieldRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type FieldAuthSendOTP struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type FieldAuthVerifyOTP struct {
	FieldAuthSendOTP
	Code string `json:"code" validate:"required"`
}

type ResponseAuthOk struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
