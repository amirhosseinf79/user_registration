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

func (o *otpRepository) CanSaveOTP(mobile string) (bool, int, error) {
	key := o.prefix + o.saveOTPLimit + mobile
	count, err := o.client.Incr(o.ctx, key).Result()
	if err != nil {
		return false, 0, err
	}
	if count == 1 {
		err = o.client.Expire(o.ctx, key, o.rateLimitDuration).Err()
		if err != nil {
			return false, 0, err
		}
	}
	if int(count) > o.otpSendRateLimit {
		return false, 0, nil
	}
	return true, o.otpSendRateLimit - int(count), nil
}

func (o *otpRepository) CanLogin(mobile string) (bool, int, error) {
	key := o.prefix + o.loginLimit + mobile
	counter, err := o.client.Incr(o.ctx, key).Result()
	if err != nil {
		return false, 0, err
	}
	if counter == 1 {
		err := o.client.Expire(o.ctx, key, o.otpExp).Err()
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
	return o.client.Set(o.ctx, o.prefix+otp.Mobile, otp.Code, o.otpExp).Err()
}

func (o *otpRepository) GetOTPByMobile(mobile string) (string, error) {
	code, err := o.client.Get(o.ctx, o.prefix+mobile).Result()
	if err == redis.Nil {
		return "", shared.ErrUsertNotFound
	}
	if err != nil {
		return "", err
	}
	return code, nil
}

func (o *otpRepository) DeleteOTP(mobile string) error {
	err := o.client.Del(o.ctx, o.prefix+mobile).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (o *otpRepository) ResetSetOTPLimit(mobile string) error {
	key := o.prefix + o.saveOTPLimit + mobile
	err := o.client.Del(o.ctx, key).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (o *otpRepository) ResetLoginLimit(mobile string) error {
	key := o.prefix + o.loginLimit + mobile
	err := o.client.Del(o.ctx, key).Err()
	if err == redis.Nil {
		return shared.ErrUsertNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
