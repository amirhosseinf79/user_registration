package model

type OTP struct {
	Prefix string `json:"prefix"`
	Key    string `json:"key"`
	Code   string `json:"code"`
}
