package main

import (
	"context"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/service"
)

func main() {
	secret := "24vm89v5y7q-x,m349ci-143-v5um120-5v27n45-1237cn4"
	gormConnStr := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tehran"
	redisAddr := "localhost:6379"
	redisPwd := ""
	debug := true

	otpTimeExp := 2 * time.Minute
	accessTokenExp := 2 * time.Hour
	refreshRokenExp := 6 * time.Hour
	smsRateLimitCount := 3
	smsRateLimitDuration := 1 * time.Minute

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

	// test
	// sendOtpFields := dto.AuthSendOTPFields{PhoneNumber: "09334429096"}
	// err := authService.SendOTP(sendOtpFields)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Code: ")
	// code, _ := reader.ReadString('\n')
	// code = code[:len(code)-2]
	// result, err := authService.VerifyOTP(dto.AuthVerifyOTPFields{AuthSendOTPFields: sendOtpFields, Code: code})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result.AccessToken, "\n", result.RefreshToken)

	// fmt.Print("refreshRoken: ")
	// rToken, _ := reader.ReadString('\n')
	// rToken = rToken[:len(rToken)-2]
	// result2, err2 := authService.RefreshToken(rToken)
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// }
	// fmt.Println(result2.AccessToken, "\n", result2.RefreshToken)
}
