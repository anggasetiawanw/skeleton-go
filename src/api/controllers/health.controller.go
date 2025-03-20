package controllers

import (
	"akatsuki/skeleton-go/src/services"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	service := services.HealthService()
	json.NewEncoder(w).Encode(service)
	return
}
