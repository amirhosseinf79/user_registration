package server

import (
	"log"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type server struct {
	validator1 interfaces.FieldValidatorMiddleware1
	access1    interfaces.AuthMiddleware1
	auth1      interfaces.AuthHandler1
	user1      interfaces.UserHandler1
	app        *fiber.App
}

func NewServer(
	authImp1 interfaces.AuthImplementation1,
) interfaces.ServerService {
	app := fiber.New()

	return &server{
		app:        app,
		validator1: authImp1.GetFieldValidator(),
		access1:    authImp1.GetAuthValidator(),
		auth1:      authImp1.GetAuthHandler(),
		user1:      authImp1.GetUserHandler(),
	}
}

func (s server) Start(port string) {
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	s.InitSwaggerRoutes()
	s.InitAuthRoutes1()
	s.InitUserRoutes1()
	err := s.app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
