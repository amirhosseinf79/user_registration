package auth

import "regexp"

func (a *authService) MatchMobile(input string) bool {
	regMobile := regexp.MustCompile(`^09\d{9}$`)
	return regMobile.MatchString(input)
}
