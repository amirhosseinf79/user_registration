package implimentation

import (
	"context"

	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	auth_handler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/auth"
	user_handler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/user"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/auth_middleware"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/field_validator"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/external"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	auth_service "github.com/amirhosseinf79/user_registration/internal/service/auth"
	"github.com/amirhosseinf79/user_registration/internal/service/email"
	"github.com/amirhosseinf79/user_registration/internal/service/jwt"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	user_service "github.com/amirhosseinf79/user_registration/internal/service/user"
)

type authImp struct {
	fieldValidator1 interfaces.FieldValidatorMiddleware1
	authValidator1  interfaces.AuthMiddleware1
	authHandler1    interfaces.AuthHandler1
	userHandler1    interfaces.UserHandler1
}

func ImplimentAuthService1(
	ctx context.Context,
	gormDB *gorm.DB,
	redisDB *redis.Client,
) interfaces.AuthImplimentation1 {
	configs := configs.NewConfig()

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

	fieldValidator1 := field_validator.NewFieldValidator()
	authValidator1 := auth_middleware.NewAuthMiddleware(jwtService)
	authHandler1 := auth_handler1.NewAuthHandler(authService)
	userHandler1 := user_handler1.NewUserHandler(userService)

	return &authImp{
		fieldValidator1: fieldValidator1,
		authValidator1:  authValidator1,
		authHandler1:    authHandler1,
		userHandler1:    userHandler1,
	}
}

func (ai *authImp) GetFieldVaidator() interfaces.FieldValidatorMiddleware1 {
	return ai.fieldValidator1
}

func (ai *authImp) GetAuthValidator() interfaces.AuthMiddleware1 {
	return ai.authValidator1
}

func (ai *authImp) GetAuthHandler() interfaces.AuthHandler1 {
	return ai.authHandler1
}

func (ai *authImp) GetUserHandler() interfaces.UserHandler1 {
	return ai.userHandler1
}
