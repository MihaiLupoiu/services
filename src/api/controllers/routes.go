package controllers

func (s *Service) initializeRoutes() {

	// Home Routes
	s.Router.HandleFunc("/", s.Home).Methods("GET")

	// Users Routes
	s.Router.HandleFunc("/user", s.CreateUser).Methods("POST")
}
