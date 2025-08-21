package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type server struct {
	fieldValidator interfaces.FieldValidatorMiddleware
	authValidator  interfaces.AuthMiddleware
	authHandler    interfaces.AuthHandler
	userHandler    interfaces.UserHandler
	app            *fiber.App
}

// @title Fiber Swagger Example API
// @version 1.0
// @description Sample API using Fiber v3 and Swagger
// @host localhost:3000
// @BasePath /api/v1
func NewServer(
	fieldValidator interfaces.FieldValidatorMiddleware,
	authValidator interfaces.AuthMiddleware,
	authHandler interfaces.AuthHandler,
	userHandler interfaces.UserHandler,
) interfaces.ServerService {
	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	return &server{
		app:            app,
		fieldValidator: fieldValidator,
		authValidator:  authValidator,
		authHandler:    authHandler,
		userHandler:    userHandler,
	}
}

func (s server) Start(port string) {
	err := s.app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
