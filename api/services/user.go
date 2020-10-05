package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"mercafacil-challenge/api/auth"
	"mercafacil-challenge/api/models"
	"mercafacil-challenge/api/responses"
	"mercafacil-challenge/api/utils"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetDB(client string, server *Server) *gorm.DB {
	if client == os.Getenv("API_SECRET_MACAPA") {
		log.Println("Using MySQL")
		return server.DBMySQL
	}
	if client == os.Getenv("API_SECRET_VAREJAO") {
		log.Println("Using PostgreSQL")
		return server.DBPostgreSQL
	}
	log.Println("Client not expected")
	return nil
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser")
	client, err := auth.ExtractClient(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("Client", client)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare(client)
	err = user.Validate(client)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	DB := GetDB(client, server)
	userCreated, err := user.CreateUser(DB)

	if err != nil {
		formattedError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers")
	client, err := auth.ExtractClient(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("Client", client)

	DB := GetDB(client, server)
	user := models.User{}
	users, err := user.FindAllUsers(DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUser")
	client, err := auth.ExtractClient(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("Client", client)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	DB := GetDB(client, server)
	user := models.User{}
	userGotten, err := user.FindUserByID(DB, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUser")
	client, err := auth.ExtractClient(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("Client", client)

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	DB := GetDB(client, server)
	user.Prepare(client)
	updatedUser, err := user.UpdateUser(DB, uint32(uid))
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteUser")
	client, err := auth.ExtractClient(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("Client", client)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	DB := GetDB(client, server)
	user := models.User{}
	_, err = user.DeleteUser(DB, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
