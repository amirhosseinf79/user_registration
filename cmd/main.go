package main

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/amirhosseinf79/user_registration/docs"
	"github.com/amirhosseinf79/user_registration/internal/application/handler"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/server"
	"github.com/amirhosseinf79/user_registration/internal/service/auth"
	"github.com/amirhosseinf79/user_registration/internal/service/jwt"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	"github.com/amirhosseinf79/user_registration/internal/service/user"
	"github.com/joho/godotenv"
)

// @title User OTP Registration API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer [...]
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// Database and secrets config
	serverPort := os.Getenv("PORT")
	secret := os.Getenv("SECRET")
	gormConnStr := os.Getenv("SQLDB")
	redisAddr := os.Getenv("RedisServer")
	redisPwd := os.Getenv("RedisPass")

	debug := false
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}

	// Expiration time configs
	otpTimeExp := 2 * time.Minute
	accessTokenExp := 2 * time.Hour
	refreshRokenExp := 6 * time.Hour
	rateLimitDuration := 10 * time.Minute
	otpSendRateLimit := 3
	loginRateLimit := 5

	ctx := context.Background()
	gormDB := database.NewGormconnection(gormConnStr, debug)
	redisDB := database.NewRedisConnection(redisAddr, redisPwd, ctx)

	otpRepo := persistence.NewOTPRepository(
		ctx,
		redisDB,
		otpTimeExp,
		loginRateLimit,
		otpSendRateLimit,
		rateLimitDuration,
	)
	tokenRepo := persistence.NewTokenRepository(ctx, redisDB, refreshRokenExp)
	userRepo := persistence.NewUserRepository(gormDB)
	jwtRepo := persistence.NewJWTRepository(secret, accessTokenExp, refreshRokenExp)

	jwtService := jwt.NewJWTService(jwtRepo, tokenRepo)
	otpService := otp.NewOTPService(otpRepo)
	userService := user.NewUserService(userRepo)

	smsService := sms.NewSMSService()
	authService := auth.NewAuthService(
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
