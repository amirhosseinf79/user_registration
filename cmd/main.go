package main

import (
	"context"

	_ "github.com/amirhosseinf79/user_registration/docs"
	auth_handler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/auth"
	user_handler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/user"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/auth_middleware"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/field_validator"
	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/external"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/server"
	auth_service "github.com/amirhosseinf79/user_registration/internal/service/auth"
	"github.com/amirhosseinf79/user_registration/internal/service/email"
	"github.com/amirhosseinf79/user_registration/internal/service/jwt"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	user_service "github.com/amirhosseinf79/user_registration/internal/service/user"
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
	mailService := email.NewEmailService()
	jwtService := jwt.NewJWTService(jwtRepo, tokenRepo)

	userService := user_service.NewUserService(
		userRepo,
		otpService,
		smsService,
		mailService,
	)

	authService := auth_service.NewAuthService(
		jwtService,
		userService,
		otpService,
		smsService,
		mailService,
	)

	fieldValidator := field_validator.NewFieldValidator()
	authValidator := auth_middleware.NewAuthMiddleware(jwtService)
	authHandler1 := auth_handler1.NewAuthHandler(authService)
	userHandler1 := user_handler1.NewUserHandler(userService)

	server := server.NewServer(
		fieldValidator,
		authValidator,
		authHandler1,
		userHandler1,
	)

	server.InitSwaggerRoutes()
	server.InitAuthRoutes1()
	server.InitUserRoutes1()
	server.Start(configs.Server.Port)
}
