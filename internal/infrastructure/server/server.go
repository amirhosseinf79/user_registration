package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	fieldValidator interfaces.FieldValidatorMiddleware
	authValidator  interfaces.AuthMiddleware
	authHandler    interfaces.AuthHandler
	userHandler    interfaces.UserHandler
	app            *fiber.App
}

func NewServer(
	fieldValidator interfaces.FieldValidatorMiddleware,
	authValidator interfaces.AuthMiddleware,
	authHandler interfaces.AuthHandler,
	userHandler interfaces.UserHandler,
) interfaces.ServerService {
	app := fiber.New()

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
