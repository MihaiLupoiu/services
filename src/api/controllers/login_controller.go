package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MihaiLupoiu/services/src/api/auth"
	"github.com/MihaiLupoiu/services/src/api/models"
	"github.com/MihaiLupoiu/services/src/api/util"
	"golang.org/x/crypto/bcrypt"
)

// Login handler
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/login
func (service *Service) Login(w http.ResponseWriter, r *http.Request) {
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

	err = user.ValidateLogin()
	if err != nil {
		util.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := service.SignIn(user.Email, user.Password)
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
	util.JSON(w, http.StatusOK, token)
}

// SignIn check that credentials are correct
func (service *Service) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}

	err = service.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
