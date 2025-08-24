package server

func (s server) InitUserRoutes() {
	user := s.app.Group("/user")
	user.Get("/all", s.userHandler.GetUsersList)
	user.Get("/:userID", s.userHandler.GetUserByID)

	profile := s.app.Group("/profile")
	profile.Put("/update", s.authValidator.CheckToken, s.fieldValidator.ValidateEmailBody, s.userHandler.UpdateProfileInfo)
}
