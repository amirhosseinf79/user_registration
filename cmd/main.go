package main

import (
	"context"

	_ "github.com/amirhosseinf79/user_registration/docs"
	"github.com/amirhosseinf79/user_registration/internal/application/handler"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware"
	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/external"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/server"
	"github.com/amirhosseinf79/user_registration/internal/service/auth"
	"github.com/amirhosseinf79/user_registration/internal/service/jwt"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	"github.com/amirhosseinf79/user_registration/internal/service/user"
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
	configs := configs.NewConfig()
	ctx := context.Background()
	gormDB := database.NewGormconnection(
		configs.DB.Gorm.ConnSTR,
		configs.Server.Debug,
	)
	redisDB := database.NewRedisConnection(configs.DB.Redis.Server, configs.DB.Redis.Password, ctx)

	otpRepo := persistence.NewOTPRepository(
		ctx,
		redisDB,
		configs.OTP.ExireTime,
		configs.OTP.LoginRate,
		configs.OTP.SendRate,
		configs.OTP.LimitDuration,
	)

	jwtRepo := persistence.NewJWTRepository(
		configs.Server.Secret,
		configs.Token.AccessTokenExp,
		configs.Token.RefreshTokenExp,
	)

	userRepo := persistence.NewUserRepository(gormDB)
	tokenRepo := persistence.NewTokenRepository(ctx, redisDB, configs.Token.RefreshTokenExp)
	smsRepo := external.NewKavenegarSMSService(configs.SMS.Kavenegar.Key, configs.SMS.Kavenegar.Sender)

	otpService := otp.NewOTPService(otpRepo)
	smsService := sms.NewSMSService(smsRepo)
	jwtService := jwt.NewJWTService(jwtRepo, tokenRepo)
	userService := user.NewUserService(userRepo, otpService, smsService)

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
	server.Start(configs.Server.Port)
}
