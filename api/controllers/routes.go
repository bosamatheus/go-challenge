package controllers

import (
	"mercafacil-challenge/api/middlewares"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) initializeRoutes() {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	s.Router.Handle("/docs", sh)
	s.Router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s.Router.HandleFunc("/api/v1/auth/jwt", s.GenerateJWT).Methods("GET")

	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.GetUserByID)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users", middlewares.MiddlewareAuthentication(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users", middlewares.MiddlewareAuthentication(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.UpdateUserByID)).Methods("PUT")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.MiddlewareAuthentication(s.DeleteUserByID)).Methods("DELETE")
}
