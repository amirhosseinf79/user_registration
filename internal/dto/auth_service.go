package dto

type AuthSendOTPFields struct {
	PhoneNumber string `json:"phoneNumber"`
}

type AuthVerifyOTPFields struct {
	AuthSendOTPFields
	Code string `json:"code"`
}

type AuthOkResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
