package configs

import "time"

type smsConfig struct {
	Key    string
	Sender string
}

type smsProviders struct {
	Kavenegar smsConfig
}

type serverConfig struct {
	Port   string
	Secret string
	Debug  bool
}

type redisDB struct {
	Server   string
	Password string
}

type gormDB struct {
	ConnSTR string
}

type dbConfig struct {
	Gorm  gormDB
	Redis redisDB
}

type tokenConfig struct {
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
}

type otpConfig struct {
	ExpireTime    time.Duration
	LimitDuration time.Duration
	SendRate      int
	LoginRate     int
}

type Configs struct {
	Server serverConfig
	DB     dbConfig
	Token  tokenConfig
	SMS    smsProviders
	OTP    otpConfig
}
