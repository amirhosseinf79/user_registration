package test

import (
	"testing"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/external"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/user_registration/internal/service/email"
	"github.com/amirhosseinf79/user_registration/internal/service/otp"
	"github.com/amirhosseinf79/user_registration/internal/service/sms"
	"github.com/amirhosseinf79/user_registration/internal/service/user"
	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB, loginMail, loginPass string) error {
	redisDB := database.NewRedisConnection("localhost:6379", "", ctx)

	userRepo := persistence.NewUserRepository(db)
	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpExp, loginRateLimit, otpSendRateLimit, rateLimitDuration)
	otpService := otp.NewOTPService(otpRepo)

	smsRepo := external.NewKavenegarSMSService("", "")
	smsService := sms.NewSMSService(smsRepo)
	mailService := email.NewEmailService()

	userService := user.NewUserService(userRepo, otpService, smsService, mailService)
	_, err := userService.RegisterUserByEmail(auth.FieldEmailRegister{
		FirstName:   "test",
		LastName:    "user",
		Email:       loginMail,
		PhoneNumber: "09334429096",
		Password:    loginPass,
	})
	if err != nil {
		return err.Error
	}
	return nil
}

func TestUserLoginOK(t *testing.T) {
	loginMail := "test@test1.com"
	loginPass := "Amir@20011"
	db := database.NewGormConnection("", true)
	userRepo := persistence.NewUserRepository(db)

	err := UserRegister(db, loginMail, loginPass)
	if err != nil {
		t.Error("should not have error", err)
	}

	userM, err2 := userRepo.GetByEmail(loginMail)
	if err2 != nil {
		t.Error("should be able to get user by email", err2)
		return
	}
	if userM.Email != loginMail {
		t.Error("should have same email")
	}
	ok := userM.ValidatePassword(loginPass)
	if !ok {
		t.Error("should be able to validate password")
	}
}

func TestUserLoginFailed(t *testing.T) {
	loginMail := "test@test2.com"
	loginPass := "Amir@20011"
	db := database.NewGormConnection("", true)
	userRepo := persistence.NewUserRepository(db)

	err := UserRegister(db, loginMail, loginPass)
	if err != nil {
		t.Error("should not have error", err)
	}

	userM, err2 := userRepo.GetByEmail(loginMail)
	if err2 != nil {
		t.Error("should be able to get user by email", err2)
		return
	}
	if userM.Email != loginMail {
		t.Error("should have same email")
	}
	ok := userM.ValidatePassword("invalidPassword")
	if ok {
		t.Error("should not be validate password")
	}
}
