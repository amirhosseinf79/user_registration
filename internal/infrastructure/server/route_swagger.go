package server

import fiberSwagger "github.com/swaggo/fiber-swagger"

func (s server) InitSwaggerRoutes() {
	s.app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
