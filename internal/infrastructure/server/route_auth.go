package server

func (s server) InitAuthRoutes1() {
	route := s.app.Group("api/v1/auth")
	route.Post("/send-otp", s.validator1.ValidateMobile, s.auth1.SendOTP)
	route.Post("/refresh-token", s.validator1.ValidateRefreshToken, s.auth1.RefreshToken)

	route.Post("/login", s.validator1.ValidateLogin, s.auth1.AutoLogin)

	registerMidd := route.Group("/register", s.validator1.ValidateRegister, s.validator1.ValidateEmail)
	registerMidd.Post("", s.validator1.ValidateMobile, s.auth1.RegisterByEmail)

	password := route.Group("/password")
	password.Post("/send-code", s.validator1.ValidateSendResetPassOTP, s.auth1.SendResetPassOTP)
	password.Put("/reset", s.validator1.ValidateVerifyCode, s.validator1.ValidateNewPassword, s.auth1.ResetPassWithOTP)
}
