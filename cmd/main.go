package main

import (
	"context"

	_ "github.com/amirhosseinf79/user_registration/docs"
	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/implementation"
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
	configF := configs.NewConfig()
	ctx := context.Background()
	gormDB := database.NewGormConnection(configF.DB.Gorm.ConnSTR, configF.Server.Debug)
	redisDB := database.NewRedisConnection(configF.DB.Redis.Server, configF.DB.Redis.Password, ctx)

	authImp1 := implementation.ImplementAuthService1(ctx, gormDB, redisDB)

	serverS := server.NewServer(authImp1)
	serverS.Start(configF.Server.Port)
}
