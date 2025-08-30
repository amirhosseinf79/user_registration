package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	fieldValidator1 interfaces.FieldValidatorMiddleware1
	authValidator1  interfaces.AuthMiddleware1
	authHandler1    interfaces.AuthHandler1
	userHandler1    interfaces.UserHandler1
	app             *fiber.App
}

func NewServer(
	authImp1 interfaces.AuthImplimentation1,
) interfaces.ServerService {
	app := fiber.New()

	return &server{
		app:             app,
		fieldValidator1: authImp1.GetFieldVaidator(),
		authValidator1:  authImp1.GetAuthValidator(),
		authHandler1:    authImp1.GetAuthHandler(),
		userHandler1:    authImp1.GetUserHandler(),
	}
}

func (s server) Start(port string) {
	err := s.app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
