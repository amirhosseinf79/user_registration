package interfaces

type ServerService interface {
	InitAuthRoutes()
	InitUserRoutes()
	Start(port string)
}
