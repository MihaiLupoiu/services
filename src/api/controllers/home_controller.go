package controllers

import (
	"net/http"

	"github.com/MihaiLupoiu/services/src/api/util"
)

// Home callback funtion
func (service *Service) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	util.JSON(w, http.StatusOK, "Users API")
}
