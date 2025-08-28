package server

func (s server) InitAuthRoutes() {
	route := s.app.Group("/auth")
	route.Post(
		"/send-otp",
		s.fieldValidator.ValidateMobile,
		s.authHandler.SendOTP,
	)
	route.Post(
		"/refresh-token",
		s.fieldValidator.ValidateRefreshToken,
		s.authHandler.RefreshToken,
	)
	route.Post(
		"/otp-login",
		s.fieldValidator.ValidateMobile,
		s.fieldValidator.ValidateCode,
		s.authHandler.LoginByOTP,
	)
	route.Post(
		"/mobile-login",
		s.fieldValidator.ValidateMobileLogin,
		s.authHandler.LoginByMobile,
	)
	route.Post(
		"/email-login",
		s.fieldValidator.ValidateEmailLogin,
		s.authHandler.LoginByEmail,
	)
	route.Post(
		"/register",
		s.fieldValidator.ValidateRegister,
		s.fieldValidator.ValidateEmailBody,
		s.fieldValidator.ValidateMobile,
		s.authHandler.RegisterByEmail,
	)
}
