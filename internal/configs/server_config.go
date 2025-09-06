package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewConfig() *Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// Database and secrets config
	ServerPort := os.Getenv("PORT")
	ServerSecret := os.Getenv("SECRET")
	GormConnStr := os.Getenv("SQLDB")
	RedisServer := os.Getenv("RedisServer")
	RedisPassword := os.Getenv("RedisPass")
	DEBUG := false
	if os.Getenv("DEBUG") == "true" {
		DEBUG = true
	}

	// external services
	KavenegarKey := os.Getenv("KAVENEGAR_KEY")
	KavenegarSender := os.Getenv("KAVENEGAR_SENDER")

	// OTP Configs
	OtpExpireTime := 2 * time.Minute
	OtpLimitDuration := 10 * time.Minute
	OtpSendRate := 3
	OtpLoginRate := 5

	// TOKEN Configs
	AccessTokenExp := 2 * time.Hour
	RefreshTokenExp := 6 * time.Hour

	return &Configs{
		Server: serverConfig{
			Port:   ServerPort,
			Secret: ServerSecret,
			Debug:  DEBUG,
		},
		DB: dbConfig{
			Gorm: gormDB{
				ConnSTR: GormConnStr,
			},
			Redis: redisDB{
				Server:   RedisServer,
				Password: RedisPassword,
			},
		},
		Token: tokenConfig{
			AccessTokenExp:  AccessTokenExp,
			RefreshTokenExp: RefreshTokenExp,
		},
		SMS: smsProviders{
			Kavenegar: smsConfig{
				Key:    KavenegarKey,
				Sender: KavenegarSender,
			},
		},
		OTP: otpConfig{
			ExpireTime:    OtpExpireTime,
			LimitDuration: OtpLimitDuration,
			SendRate:      OtpSendRate,
			LoginRate:     OtpLoginRate,
		},
	}
}
