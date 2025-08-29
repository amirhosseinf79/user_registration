package interfaces

type ServerService interface {
	InitSwaggerRoutes()
	InitAuthRoutes1()
	InitUserRoutes1()
	Start(port string)
}
