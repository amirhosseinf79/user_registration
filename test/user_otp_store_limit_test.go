package test

import (
	"context"
	"testing"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
)

var (
	ctx               = context.Background()
	otpExp            = 1 * time.Minute
	loginRateLimit    = 5
	otpSendRateLimit  = 3
	rateLimitDuration = 3 * time.Minute
)

func TestUserOTPStoreLimited(t *testing.T) {
	redisDB := database.NewRedisConnection("localhost:6379", "", ctx)
	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpExp, loginRateLimit, otpSendRateLimit, rateLimitDuration)

	var canLogin bool
	var err error
	for range otpSendRateLimit + 1 {
		canLogin, _, err = otpRepo.CanSaveOTP("smst:09334429096")
	}

	if err != nil {
		t.Error("should not have error", err)
	}

	if canLogin {
		t.Error("otpService should not be able to store OTP")
	}
}

func TestUserOTPStoreNotLimited(t *testing.T) {
	redisDB := database.NewRedisConnection("localhost:6379", "", ctx)
	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpExp, loginRateLimit, otpSendRateLimit, rateLimitDuration)

	var canLogin bool
	var err error
	for range otpSendRateLimit {
		canLogin, _, err = otpRepo.CanSaveOTP("smst:09334429097")
	}

	if err != nil {
		t.Error("should not have error", err)
	}

	if !canLogin {
		t.Error("otpService should be able to store OTP")
	}
}
