package services

import "mercafacil-challenge/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", s.Swagger).Methods("GET")
	s.Router.HandleFunc("/swagger/*", s.Swagger).Methods("GET")

	s.Router.HandleFunc("/api/v1/auth/jwt", s.GenerateJWT).Methods("GET")

	s.Router.HandleFunc("/api/v1/users", middlewares.MiddlewareAuthentication(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/users", middlewares.MiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
}
