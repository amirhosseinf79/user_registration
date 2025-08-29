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
		"/login",
		s.fieldValidator.ValidateLogin,
		s.authHandler.AutoLogin,
	)

	route.Post(
		"/register",
		s.fieldValidator.ValidateRegister,
		s.fieldValidator.ValidateEmail,
		s.fieldValidator.ValidateMobile,
		s.authHandler.RegisterByEmail,
	)

	password := route.Group("/password")
	password.Post(
		"/send-code",
		s.fieldValidator.ValidateSendResetPassOTP,
		s.authHandler.SendResetPassOTP,
	)
	password.Put(
		"/reset",
		s.fieldValidator.ValidateVerifyCode,
		s.fieldValidator.ValidateNewPassword,
		s.authHandler.ResetPassWithOTP,
	)
}
