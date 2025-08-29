package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	fieldValidator interfaces.FieldValidatorMiddleware
	authValidator  interfaces.AuthMiddleware
	authHandler1   interfaces.AuthHandler1
	userHandler1   interfaces.UserHandler1
	app            *fiber.App
}

func NewServer(
	fieldValidator interfaces.FieldValidatorMiddleware,
	authValidator interfaces.AuthMiddleware,
	authHandler1 interfaces.AuthHandler1,
	userHandler1 interfaces.UserHandler1,
) interfaces.ServerService {
	app := fiber.New()

	return &server{
		app:            app,
		fieldValidator: fieldValidator,
		authValidator:  authValidator,
		authHandler1:   authHandler1,
		userHandler1:   userHandler1,
	}
}

func (s server) Start(port string) {
	err := s.app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
