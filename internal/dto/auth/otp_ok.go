package auth

import "time"

type OTPOk struct {
	Code       int           `json:"code,omitempty"`
	ExpiresIn  time.Duration `json:"expiresIn"`
	RetryCount int           `json:"retryCount"`
}

type OTPOkMock struct {
	Code       int `json:"code,omitempty"`
	ExpiresIn  int `json:"expiresIn"`
	RetryCount int `json:"retryCount"`
}
