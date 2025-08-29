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
	login := route.Group("/login")
	login.Post(
		"/otp",
		s.fieldValidator.ValidateMobile,
		s.fieldValidator.ValidateVerifyField,
		s.authHandler.LoginByOTP,
	)
	login.Post(
		"/mobile",
		s.fieldValidator.ValidateMobileLogin,
		s.authHandler.LoginByMobile,
	)
	login.Post(
		"/email",
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
