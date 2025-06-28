package controller

func (s *Server) SetupRoutes() {
	s.Router.GET("/", s.GetHealth)
}