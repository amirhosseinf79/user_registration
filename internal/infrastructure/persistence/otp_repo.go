package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type otpRepository struct {
	ctx    context.Context
	client *redis.Client
	prefix string
}

func NewOTPRepository(ctx context.Context, client *redis.Client) repository.OTPRepository {
	return &otpRepository{
		prefix: "OTP:",
		client: client,
		ctx:    ctx,
	}
}

func (o *otpRepository) Set(otp *model.OTP, exp time.Duration) error {
	return o.client.Set(o.ctx, o.prefix+otp.Mobile, otp, exp).Err()
}

func (o *otpRepository) Get(otp *model.OTP) error {
	errStr := o.client.HGetAll(o.ctx, o.prefix+otp.Mobile).Scan(otp).Error()
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
