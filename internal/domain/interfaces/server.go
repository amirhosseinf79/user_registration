package interfaces

type ServerService interface {
	InitAuthRoutes()
	Start(port string)
}
