package server

func (s server) InitUserRoutes() {
	user := s.app.Group("/user")
	user.Get("/all", s.userHandler.GetUsersList)
	user.Get("/:userID", s.userHandler.GetUserByID)

	profile := s.app.Group("/profile", s.authValidator.CheckToken)
	profile.Get("/", s.userHandler.GetUserProfile)
	profile.Post("/send/otp", s.userHandler.SendUserOTP)
	profile.Patch("/update/info", s.fieldValidator.ValidateEmail, s.userHandler.UpdateProfileInfo)
	profile.Put("/update/password", s.fieldValidator.ValidateNewPassword, s.userHandler.UpdateUserPassword)
	profile.Post("/verify/mobile", s.fieldValidator.ValidateVerifyCode, s.userHandler.VerifyUserOTP)
}
