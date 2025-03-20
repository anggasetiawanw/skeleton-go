package api

import (
	"akatsuki/skeleton-go/src/api/controllers"

	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", controllers.Health)
	router.GET("/health", controllers.Health)

	router.GET("/AddDevice", controllers.AddDevice)

	return router
}
