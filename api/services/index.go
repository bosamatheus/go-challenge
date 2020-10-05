package services

import (
	"net/http"
)

func (server *Server) Swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "./docs/swagger.json")
}
