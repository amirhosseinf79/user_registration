package auth_request

type FieldRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
