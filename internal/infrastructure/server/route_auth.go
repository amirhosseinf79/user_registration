package server

func (s server) InitAuthRoutes1() {
	route := s.app.Group("api/v1/auth")
	route.Post(
		"/send-otp",
		s.fieldValidator1.ValidateMobile,
		s.authHandler1.SendOTP,
	)
	route.Post(
		"/refresh-token",
		s.fieldValidator1.ValidateRefreshToken,
		s.authHandler1.RefreshToken,
	)

	route.Post(
		"/login",
		s.fieldValidator1.ValidateLogin,
		s.authHandler1.AutoLogin,
	)

	route.Post(
		"/register",
		s.fieldValidator1.ValidateRegister,
		s.fieldValidator1.ValidateEmail,
		s.fieldValidator1.ValidateMobile,
		s.authHandler1.RegisterByEmail,
	)

	password := route.Group("/password")
	password.Post(
		"/send-code",
		s.fieldValidator1.ValidateSendResetPassOTP,
		s.authHandler1.SendResetPassOTP,
	)
	password.Put(
		"/reset",
		s.fieldValidator1.ValidateVerifyCode,
		s.fieldValidator1.ValidateNewPassword,
		s.authHandler1.ResetPassWithOTP,
	)
}
