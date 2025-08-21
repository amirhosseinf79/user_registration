package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
)

type server struct {
	fieldValidator interfaces.FieldValidatorMiddleware
	authHandler    interfaces.AuthHandler
	app            *fiber.App
}

func NewServer(
	fieldValidator interfaces.FieldValidatorMiddleware,
	authHandler interfaces.AuthHandler,
) interfaces.ServerService {
	app := fiber.New()
	return &server{
		app:            app,
		fieldValidator: fieldValidator,
		authHandler:    authHandler,
	}
}

func (s server) Start(port string) {
	err := s.app.Listen(":"+port, fiber.ListenConfig{EnablePrefork: false})
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
