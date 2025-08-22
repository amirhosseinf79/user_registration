package server

func (s server) InitAuthRoutes() {
	route := s.app.Group("/auth")
	route.Post("/send-otp", s.fieldValidator.ValidateMobile, s.authHandler.SendOTP)
	route.Post("/verify-otp", s.fieldValidator.ValidateMobile, s.fieldValidator.ValidateCode, s.authHandler.LoginByOTP)
	route.Post("/refresh-token", s.fieldValidator.ValidateRefreshToken, s.authHandler.RefreshToken)
}
