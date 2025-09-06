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
	SERVER_PORT := os.Getenv("PORT")
	SERVER_SECRET := os.Getenv("SECRET")
	GORM_CONNSTR := os.Getenv("SQLDB")
	REDIS_SERVER := os.Getenv("RedisServer")
	REDIS_PASSWORD := os.Getenv("RedisPass")
	DEBUG := false
	if os.Getenv("DEBUG") == "true" {
		DEBUG = true
	}

	// external services
	KAVENEGAR_KEY := os.Getenv("KAVENEGAR_KEY")
	KAVENEGAR_SENDER := os.Getenv("KAVENEGAR_SENDER")

	// OTP Configs
	OTP_EXPIRE_TIME := 2 * time.Minute
	OTP_LIMIT_DURATION := 10 * time.Minute
	OTP_SEND_RATE := 3
	OTP_LOGIN_RATE := 5

	// TOKEN Configs
	ACCESS_TOKEN_EXP := 2 * time.Hour
	REFRESH_TOKEN_EXP := 6 * time.Hour

	return &Configs{
		Server: serverConfig{
			Port:   SERVER_PORT,
			Secret: SERVER_SECRET,
			Debug:  DEBUG,
		},
		DB: dbConfig{
			Gorm: gormDB{
				ConnSTR: GORM_CONNSTR,
			},
			Redis: redisDB{
				Server:   REDIS_SERVER,
				Password: REDIS_PASSWORD,
			},
		},
		Token: tokenConfig{
			AccessTokenExp:  ACCESS_TOKEN_EXP,
			RefreshTokenExp: REFRESH_TOKEN_EXP,
		},
		SMS: smsProviders{
			Kavenegar: smsConfig{
				Key:    KAVENEGAR_KEY,
				Sender: KAVENEGAR_SENDER,
			},
		},
		OTP: otpConfig{
			ExpireTime:    OTP_EXPIRE_TIME,
			LimitDuration: OTP_LIMIT_DURATION,
			SendRate:      OTP_SEND_RATE,
			LoginRate:     OTP_LOGIN_RATE,
		},
	}
}
