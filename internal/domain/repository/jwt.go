package repository

type JWTRepository interface {
	GenerateToken(userID uint, long bool) (string, error)
	Verify(tokenString string) (uint, error)
}
