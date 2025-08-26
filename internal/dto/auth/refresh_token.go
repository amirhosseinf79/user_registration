package auth

type FieldRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
