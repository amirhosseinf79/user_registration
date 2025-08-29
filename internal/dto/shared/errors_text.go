package shared

import "errors"

var ErrUsertNotFound = errors.New("user not found")
var ErrInvalidToken = errors.New("invalid token")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidUser = errors.New("invalid user")
var ErrInvalidCreds = errors.New("invalid credentials")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidIssuer = errors.New("invalid issuer")
var ErrTokenExpired = errors.New("token expired")
var ErrInvalidCode = errors.New("invalid code")
var ErrInvalidMobile = errors.New("invalid phone number")
var ErrInvalidEmail = errors.New("invalid email")
var ErrEmailExists = errors.New("email already exists")
var ErrMobileExists = errors.New("mobile already exists")
var ErrSmsRateLimited = errors.New("can not send sms")
var ErrInternalServerError = errors.New("internal server error")
var ErrPasswordValidation = errors.New("The password must be at least 10 characters long and include at least one special character.")
