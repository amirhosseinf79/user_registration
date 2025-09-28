package test

import (
	"testing"

	"github.com/amirhosseinf79/user_registration/internal/infrastructure/database"
	"github.com/amirhosseinf79/user_registration/internal/infrastructure/persistence"
)

func TestUserLoginLimited(t *testing.T) {
	redisDB := database.NewRedisConnection("localhost:6379", "", ctx)
	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpExp, loginRateLimit, otpSendRateLimit, rateLimitDuration)

	var canStore bool
	var err error
	for range loginRateLimit + 1 {
		canStore, _, err = otpRepo.CanLogin("09334429096")
	}

	if err != nil {
		t.Error("should not have error", err)
	}

	if canStore {
		t.Error("otpService limit: should not be able to login")
	}
}

func TestUserLoginNotLimited(t *testing.T) {
	redisDB := database.NewRedisConnection("localhost:6379", "", ctx)
	otpRepo := persistence.NewOTPRepository(ctx, redisDB, otpExp, loginRateLimit, otpSendRateLimit, rateLimitDuration)

	var canStore bool
	var err error
	for range loginRateLimit {
		canStore, _, err = otpRepo.CanLogin("09334429097")
	}

	if err != nil {
		t.Error("should not have error", err)
	}

	if !canStore {
		t.Error("otpService limit: should be able to login")
	}
}
