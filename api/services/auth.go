package services

import (
	"net/http"

	"mercafacil-challenge/api/auth"
	"mercafacil-challenge/api/responses"
	"mercafacil-challenge/api/utils"
)

func (server *Server) GenerateJWT(w http.ResponseWriter, r *http.Request) {
	client := r.Header.Get("client")
	if client != "" {
		token, err := auth.CreateToken(client)
		if err != nil {
			formattedError := utils.FormatError(err.Error())
			responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
			return
		}
		responses.JSON(w, http.StatusOK, token)
	} else {
		responses.JSON(w, http.StatusBadRequest, nil)
	}
}
