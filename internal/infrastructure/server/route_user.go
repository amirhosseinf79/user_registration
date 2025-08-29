package server

func (s server) InitUserRoutes() {
	user := s.app.Group("/user")
	user.Get("/all", s.userHandler.GetUsersList)
	user.Get("/:userID", s.userHandler.GetUserByID)

	profile := s.app.Group("/profile", s.authValidator.CheckToken)
	profile.Get("/", s.userHandler.GetUserProfile)
	profile.Patch("/update/info", s.fieldValidator.ValidateEmail, s.userHandler.UpdateProfileInfo)
	profile.Put("/update/password", s.fieldValidator.ValidateNewPassword, s.userHandler.UpdateUserPassword)

	profile.Post("/send/verify-mobile-otp", s.userHandler.SendUserVerifyMobile)
	profile.Post("/send/verify-email-otp", s.userHandler.SendUserVerifyEmail)
	profile.Post("/verify/mobile", s.fieldValidator.ValidateVerifyCode, s.userHandler.VerifyUserMobile)
	profile.Post("/verify/email", s.fieldValidator.ValidateVerifyCode, s.userHandler.VerifyUserEmail)
}
