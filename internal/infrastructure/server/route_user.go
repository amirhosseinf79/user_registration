package server

func (s server) InitUserRoutes1() {
	user := s.app.Group("api/v1/user")
	user.Get("/all", s.userHandler1.GetUsersList)
	user.Get("/:userID", s.userHandler1.GetUserByID)

	profile := s.app.Group("/profile", s.authValidator.CheckToken)
	profile.Get("/", s.userHandler1.GetUserProfile)
	profile.Patch("/update/info", s.fieldValidator.ValidateEmail, s.userHandler1.UpdateProfileInfo)
	profile.Put("/update/password", s.fieldValidator.ValidateNewPassword, s.userHandler1.UpdateUserPassword)

	profile.Post("/send/verify-mobile-otp", s.userHandler1.SendUserVerifyMobile)
	profile.Post("/send/verify-email-otp", s.userHandler1.SendUserVerifyEmail)
	profile.Post("/verify/mobile", s.fieldValidator.ValidateVerifyCode, s.userHandler1.VerifyUserMobile)
	profile.Post("/verify/email", s.fieldValidator.ValidateVerifyCode, s.userHandler1.VerifyUserEmail)
}
