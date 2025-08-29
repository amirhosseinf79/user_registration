package server

func (s server) InitAuthRoutes1() {
	route := s.app.Group("api/v1/auth")
	route.Post(
		"/send-otp",
		s.fieldValidator.ValidateMobile,
		s.authHandler1.SendOTP,
	)
	route.Post(
		"/refresh-token",
		s.fieldValidator.ValidateRefreshToken,
		s.authHandler1.RefreshToken,
	)

	route.Post(
		"/login",
		s.fieldValidator.ValidateLogin,
		s.authHandler1.AutoLogin,
	)

	route.Post(
		"/register",
		s.fieldValidator.ValidateRegister,
		s.fieldValidator.ValidateEmail,
		s.fieldValidator.ValidateMobile,
		s.authHandler1.RegisterByEmail,
	)

	password := route.Group("/password")
	password.Post(
		"/send-code",
		s.fieldValidator.ValidateSendResetPassOTP,
		s.authHandler1.SendResetPassOTP,
	)
	password.Put(
		"/reset",
		s.fieldValidator.ValidateVerifyCode,
		s.fieldValidator.ValidateNewPassword,
		s.authHandler1.ResetPassWithOTP,
	)
}
