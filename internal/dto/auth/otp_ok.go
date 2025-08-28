package auth

import "time"

type OTPOk struct {
	Code       int           `json:"code,omitempty"`
	TTL        time.Duration `json:"ttl"`
	RetryCount int           `json:"retryCount"`
}

type OTPOkMock struct {
	Code       int `json:"code,omitempty"`
	TTL        int `json:"ttl"`
	RetryCount int `json:"retryCount"`
}
