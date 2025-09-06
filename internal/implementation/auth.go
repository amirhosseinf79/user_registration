package implementation

import (
	"context"

	"github.com/amirhosseinf79/user_registration/internal/configs"
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	authhandler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/auth"
	userhandler1 "github.com/amirhosseinf79/user_registration/internal/application/handler/v1/user"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/auth_middleware"
	"github.com/amirhosseinf79/user_registration/internal/application/middleware/v1/field_validator"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/external"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	authservice "github.com/amirhosseinf79/user_registration/internal/service/auth"
	"github.com/amirhosseinf79/user_registration/internal/service/email"
	"github.com/amirhosseinf79/user_registration/internal/service/jwt"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	userservice "github.com/amirhosseinf79/user_registration/internal/service/user"
)

type authImp struct {
	fieldValidator1 interfaces.FieldValidatorMiddleware1
	authValidator1  interfaces.AuthMiddleware1
	authHandler1    interfaces.AuthHandler1
	userHandler1    interfaces.UserHandler1
}

func ImplementAuthService1(
	ctx context.Context,
	gormDB *gorm.DB,
	redisDB *redis.Client,
) interfaces.AuthImplementation1 {
	config := configs.NewConfig()

	otpRepo := persistence.NewOTPRepository(
		ctx,
		redisDB,
		config.OTP.ExpireTime,
		config.OTP.LoginRate,
		config.OTP.SendRate,
		config.OTP.LimitDuration,
	)

	jwtRepo := persistence.NewJWTRepository(
		config.Server.Secret,
		config.Token.AccessTokenExp,
		config.Token.RefreshTokenExp,
	)

	userRepo := persistence.NewUserRepository(gormDB)
	tokenRepo := persistence.NewTokenRepository(ctx, redisDB, config.Token.RefreshTokenExp)
	smsRepo := external.NewKavenegarSMSService(config.SMS.Kavenegar.Key, config.SMS.Kavenegar.Sender)

	otpService := otp.NewOTPService(otpRepo)
	smsService := sms.NewSMSService(smsRepo)
	mailService := email.NewEmailService()
	jwtService := jwt.NewJWTService(jwtRepo, tokenRepo)

	userService := userservice.NewUserService(
		userRepo,
		otpService,
		smsService,
		mailService,
	)

	authService := authservice.NewAuthService(
		jwtService,
		userService,
		otpService,
		smsService,
		mailService,
	)

	fieldValidator1 := field_validator.NewFieldValidator()
	authValidator1 := auth_middleware.NewAuthMiddleware(jwtService)
	authHandler1 := authhandler1.NewAuthHandler(authService)
	userHandler1 := userhandler1.NewUserHandler(userService)

	return &authImp{
		fieldValidator1: fieldValidator1,
		authValidator1:  authValidator1,
		authHandler1:    authHandler1,
		userHandler1:    userHandler1,
	}
}

func (ai *authImp) GetFieldValidator() interfaces.FieldValidatorMiddleware1 {
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
