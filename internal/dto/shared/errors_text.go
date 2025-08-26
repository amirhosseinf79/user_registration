package shared

import "errors"

var ErrUsertNotFound = errors.New("user not found")
var ErrInvalidToken = errors.New("invalid token")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidUser = errors.New("invalid user")
var ErrInvalidIssuer = errors.New("invalid issuer")
var ErrTokenExpired = errors.New("token expired")
var ErrInvalidCode = errors.New("invalid code")
var ErrInvalidMobile = errors.New("invalid phone number")
var ErrInvalidEmail = errors.New("invalid email")
var ErrSmsRateLimited = errors.New("can not send sms")
var ErrInternalServerError = errors.New("internal server error")
