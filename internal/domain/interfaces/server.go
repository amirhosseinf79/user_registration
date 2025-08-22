package interfaces

type ServerService interface {
	InitSwaggerRoutes()
	InitAuthRoutes()
	InitUserRoutes()
	Start(port string)
}
