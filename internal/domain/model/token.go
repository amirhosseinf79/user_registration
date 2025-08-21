package model

type Token struct {
	UserID       uint
	RefreshToken string `json:"refreshToken"`
}
