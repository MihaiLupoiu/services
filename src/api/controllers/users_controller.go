package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/MihaiLupoiu/services/src/api/models"
	"github.com/MihaiLupoiu/services/src/api/util"
)

// CreateUser callback funtion
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/user
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

	userCreated, err := user.Create(service.DB)
	if err != nil {
		errorMessage := errors.New("Incorrect Details")

		if strings.Contains(err.Error(), "email") {
			errorMessage = errors.New("Email Already Taken")
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
