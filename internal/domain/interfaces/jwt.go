package interfaces

type JWTInterface interface {
	GenerateToken(userID uint, long bool) (string, error)
	Verify(tokenString string) (uint, error)
}
