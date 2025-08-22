package main

import (
	"context"
	"time"

	_ "github.com/amirhosseinf79/user_registration/docs"
	"github.com/amirhosseinf79/user_registration/internal/application/handler"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/server"
	"github.com/amirhosseinf79/user_registration/internal/service"
)

// @title User OTP Registration API
// @version 1.0
// @description User OTP Registration API
// @BasePath /
// @schemes http
func main() {
	// serverPort := os.Getenv("PORT")
	// secret := os.Getenv("SECRET")
	// gormConnStr := os.Getenv("SQLDB")
	// redisAddr := os.Getenv("RedisServer")
	// redisPwd := os.Getenv("RedisPass")

	// debug := false
	// if os.Getenv("DEBUG") == "true" {
	// 	debug = true
	// }
	serverPort := "8080"
	secret := "24vm89v5y7q-x,m349ci-143-v5um120-5v27n45-1237cn4"
	gormConnStr := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tehran"
	redisAddr := "localhost:6379"
	redisPwd := ""
	debug := true

	otpTimeExp := 2 * time.Minute
	accessTokenExp := 2 * time.Hour
	refreshRokenExp := 6 * time.Hour
	smsRateLimitDuration := 1 * time.Minute
	smsRateLimitCount := 3

	ctx := context.Background()
	gormDB := database.NewGormconnection(gormConnStr, debug)
	redisDB := database.NewRedisConnection(redisAddr, redisPwd, ctx)

	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpTimeExp, smsRateLimitCount, smsRateLimitDuration)
	tokenRepo := persistence.NewTokenRepository(ctx, redisDB, refreshRokenExp)
	userRepo := persistence.NewUserRepository(gormDB)
	jwtRepo := persistence.NewJWTRepository(secret, accessTokenExp, refreshRokenExp)

	jwtService := service.NewJWTService(jwtRepo, tokenRepo)
	otpService := service.NewOTPService(otpRepo)
	userService := service.NewUserService(userRepo)

	smsService := service.NewSMSService()
	authService := service.NewAuthService(
		jwtService,
		userService,
		otpService,
		smsService,
	)

	fieldValidator := middleware.NewFieldValidator()
	authValidator := middleware.NewAuthMiddleware(jwtService)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	server := server.NewServer(
		fieldValidator,
		authValidator,
		authHandler,
		userHandler,
	)

	server.InitSwaggerRoutes()
	server.InitAuthRoutes()
	server.InitUserRoutes()
	server.Start(serverPort)
}
