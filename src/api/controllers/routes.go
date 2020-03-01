package controllers

func (s *Service) initializeRoutes() {

	// Home Routes
	s.Router.HandleFunc("/", s.Home).Methods("GET")

	// Users Routes
	s.Router.HandleFunc("/user/add", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/user/{id}", s.GetUserByID).Methods("GET")
	s.Router.HandleFunc("/user/{id}", s.UpdateUser).Methods("POST")
	s.Router.HandleFunc("/user/{id}", s.DeleteUser).Methods("DELETE")
}
