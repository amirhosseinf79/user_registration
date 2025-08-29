package persistence

import (
	"context"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/redis/go-redis/v9"
)

type otpRepository struct {
	prefix            string
	loginLimit        string
	saveOTPLimit      string
	ctx               context.Context
	client            *redis.Client
	otpExp            time.Duration
	loginRateLimit    int
	otpSendRateLimit  int
	rateLimitDuration time.Duration
}

func NewOTPRepository(
	ctx context.Context,
	client *redis.Client,
	otpExp time.Duration,
	loginRateLimit int,
	otpSendRateLimit int,
	rateLimitDuration time.Duration,
) repository.OTPRepository {
	return &otpRepository{
		prefix:            "OTP:",
		loginLimit:        "login:limit:",
		saveOTPLimit:      "store:limit:",
		client:            client,
		ctx:               ctx,
		otpExp:            otpExp,
		loginRateLimit:    loginRateLimit,
		otpSendRateLimit:  otpSendRateLimit,
		rateLimitDuration: rateLimitDuration,
	}
}

func (o *otpRepository) GetOTPExpDuration() time.Duration {
	return o.otpExp / 1000000000
}

func (o *otpRepository) CanSaveOTP(key string) (bool, int, error) {
	k := o.prefix + o.saveOTPLimit + key
	count, err := o.client.Incr(o.ctx, k).Result()
	if err != nil {
		return false, 0, err
	}
	if count == 1 {
		err = o.client.Expire(o.ctx, k, o.rateLimitDuration).Err()
		if err != nil {
			return false, 0, err
		}
	}
	if int(count) > o.otpSendRateLimit {
		return false, 0, nil
	}
	return true, o.otpSendRateLimit - int(count), nil
}

func (o *otpRepository) CanLogin(key string) (bool, int, error) {
	k := o.prefix + o.loginLimit + key
	counter, err := o.client.Incr(o.ctx, k).Result()
	if err != nil {
		return false, 0, err
	}
	if counter == 1 {
		err := o.client.Expire(o.ctx, k, o.otpExp).Err()
		if err != nil {
			return false, 0, err
		}
	}
	if int(counter) > o.loginRateLimit {
		return false, 0, nil
	}
	return true, o.loginRateLimit - int(counter), nil
}

func (o *otpRepository) SaveOTP(otp *model.OTP) error {
	return o.client.Set(o.ctx, o.prefix+otp.Prefix+otp.Key, otp.Code, o.otpExp).Err()
}

func (o *otpRepository) GetOTP(prefix, key string) (string, error) {
	code, err := o.client.Get(o.ctx, o.prefix+prefix+key).Result()
	if err == redis.Nil {
		return "", shared.ErrUsertNotFound
	}
	if err != nil {
		return "", err
	}
	return code, nil
}

func (o *otpRepository) DeleteOTP(prefix, key string) error {
	err := o.client.Del(o.ctx, o.prefix+prefix+key).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (o *otpRepository) ResetSetOTPLimit(key string) error {
	k := o.prefix + o.saveOTPLimit + key
	err := o.client.Del(o.ctx, k).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (o *otpRepository) ResetLoginLimit(key string) error {
	k := o.prefix + o.loginLimit + key
	err := o.client.Del(o.ctx, k).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
