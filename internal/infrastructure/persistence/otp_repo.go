package persistence

import (
	"context"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/redis/go-redis/v9"
)

type otpRepository struct {
	ctx               context.Context
	client            *redis.Client
	prefix            string
	otpExp            time.Duration
	rateLimitCount    int
	rateLimitDuration time.Duration
}

func NewOTPRepository(
	ctx context.Context,
	client *redis.Client,
	otpExp time.Duration,
	rateLimitCount int,
	rateLimitDuration time.Duration,
) repository.OTPRepository {
	return &otpRepository{
		prefix:            "OTP:",
		client:            client,
		ctx:               ctx,
		otpExp:            otpExp,
		rateLimitCount:    rateLimitCount,
		rateLimitDuration: rateLimitDuration,
	}
}

func (o *otpRepository) CanSetOTP(mobile string) (bool, error) {
	key := o.prefix + "limit:" + mobile
	count, err := o.client.Incr(o.ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		err = o.client.Expire(o.ctx, key, o.rateLimitDuration).Err()
		if err != nil {
			return false, err
		}
	}
	if int(count) > o.rateLimitCount {
		return false, nil
	}
	return true, nil
}

func (o *otpRepository) SaveOTP(otp *model.OTP) error {
	return o.client.Set(o.ctx, o.prefix+otp.Mobile, otp.Code, o.otpExp).Err()
}

func (o *otpRepository) GetOTPByMobile(mobile string) (string, error) {
	code, err := o.client.Get(o.ctx, o.prefix+mobile).Result()
	if err == redis.Nil {
		return "", dto.ErrObjectNotFound
	}
	if err != nil {
		return "", err
	}
	return code, nil
}

func (o *otpRepository) DeleteOTP(mobile string) error {
	err := o.client.Del(o.ctx, o.prefix+mobile).Err()
	if err == redis.Nil {
		return dto.ErrObjectNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
