package auth

type ResponseJWT struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewJWTResponse(accessToken, refreshToken string) *ResponseJWT {
	return &ResponseJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
