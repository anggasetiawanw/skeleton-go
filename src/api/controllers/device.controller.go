package controllers

import (
	"akatsuki/skeleton-go/src/services"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AddDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	service, _ := services.GetAllDevice()
	json.NewEncoder(w).Encode(service)
	return
}
