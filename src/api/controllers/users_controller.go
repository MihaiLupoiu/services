package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/MihaiLupoiu/services/src/api/models"
	"github.com/MihaiLupoiu/services/src/api/util"
)

// CreateUser callback funtion
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/user/add
func (service *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = json.Unmarshal(body, &user)
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Validate()
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.HashPassword()
	userCreated, err := user.Create(service.DB)
	if err != nil {
		errorMessage := errors.New("Incorrect Use")

		if strings.Contains(err.Error(), "email") {
			errorMessage = errors.New("Email Already in Use")
		}
		if strings.Contains(err.Error(), "hashedPassword") {
			errorMessage = errors.New("Incorrect Password")
		}

		util.JSONError(w, http.StatusInternalServerError, errorMessage)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	util.JSON(w, http.StatusCreated, userCreated)
}

// GetUserByID handler
// curl -i -X GET http://localhost:8080/user/1
func (service *Service) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if authenticated
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindByID(service.DB, uint32(uid))
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	util.JSON(w, http.StatusOK, userGotten)
}

// UpdateUser handler
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"email\": \"jd@fakeee.com\",\"password\": \"1234\"}" http://localhost:8080/user/1
func (service *Service) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Autenticate
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate()
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.Update(service.DB, uint32(uid))
	if err != nil {
		errorMessage := errors.New("Incorrect Use")

		if strings.Contains(err.Error(), "email") {
			errorMessage = errors.New("Email Already in Use")
		}
		if strings.Contains(err.Error(), "hashedPassword") {
			errorMessage = errors.New("Incorrect Password")
		}

		util.JSONError(w, http.StatusInternalServerError, errorMessage)
		return
	}
	util.JSON(w, http.StatusOK, updatedUser)
}

// DeleteUser handler
// curl -i -X DELETE http://localhost:8080/user/1
func (service *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if authenticated

	vars := mux.Vars(r)
	user := models.User{}
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}

	_, err = user.Delete(service.DB, uint32(uid))
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	util.JSON(w, http.StatusNoContent, "")
}
