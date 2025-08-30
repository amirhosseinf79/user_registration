package server

func (s server) InitUserRoutes1() {
	user := s.app.Group("api/v1/user")
	user.Get("/all", s.user1.GetUsersList)
	user.Get("/:userID", s.user1.GetUserByID)

	profile := s.app.Group("api/v1/profile", s.access1.CheckToken)
	profile.Get("/", s.user1.GetUserProfile)
	profile.Patch("/update/info", s.validator1.ValidateEmail, s.validator1.ValidateUpdateMobile, s.user1.UpdateProfileInfo)
	profile.Put("/update/password", s.validator1.ValidateNewPassword, s.user1.UpdateUserPassword)

	profile.Post("/send/verify-mobile-otp", s.user1.SendUserVerifyMobile)
	profile.Post("/send/verify-email-otp", s.user1.SendUserVerifyEmail)
	profile.Post("/verify/mobile", s.validator1.ValidateVerifyCode, s.user1.VerifyUserMobile)
	profile.Post("/verify/email", s.validator1.ValidateVerifyCode, s.user1.VerifyUserEmail)
}
