package server

func (s server) InitUserRoutes() {
	user := s.app.Group("/user")
	user.Get("/all", s.userHandler.GetUsersList)
	user.Get("/:userID", s.userHandler.GetUserByID)

	profile := s.app.Group("/profile", s.authValidator.CheckToken)
	profile.Get("/", s.userHandler.GetUserProfile)
	profile.Patch("/update/info", s.fieldValidator.ValidateEmailBody, s.userHandler.UpdateProfileInfo)
	profile.Put("/update/password", s.fieldValidator.ValidateNewPassword, s.userHandler.UpdateUserPassword)
	profile.Post("/verify/mobile", s.fieldValidator.ValidateVerifyCode, s.userHandler.VerifyUserOTP)
}
