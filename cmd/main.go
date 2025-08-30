package main

import (
	"context"

	_ "github.com/amirhosseinf79/user_registration/docs"
	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/implimentation"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/server"
)

// @title User OTP Registration API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer [...]
// @schemes http
// @BasePath /
func main() {
	configs := configs.NewConfig()
	ctx := context.Background()
	gormDB := database.NewGormconnection(configs.DB.Gorm.ConnSTR, configs.Server.Debug)
	redisDB := database.NewRedisConnection(configs.DB.Redis.Server, configs.DB.Redis.Password, ctx)

	authImp1 := implimentation.ImplimentAuthService1(ctx, gormDB, redisDB)

	server := server.NewServer(authImp1)
	server.InitSwaggerRoutes()
	server.InitAuthRoutes1()
	server.InitUserRoutes1()
	server.Start(configs.Server.Port)
}
