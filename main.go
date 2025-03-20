package main

import (
	"akatsuki/skeleton-go/src/api"
	"akatsuki/skeleton-go/src/api/middlewares"
	"akatsuki/skeleton-go/src/config"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)
func main() {
	cfg := config.NewConfig()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
	})
	
	env := cfg.GlobalEnvironment(config.Envs{})

	s := &http.Server{
		Addr:    ":" + env.Port,
		Handler: c.Handler(middlewares.Logging(api.Router())),
		WriteTimeout: time.Duration(env.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(env.ReadTimeout) * time.Second,
	}

	log.Printf("Service Started on Port %s", s.Addr)
	logrus.Debug(s.ListenAndServe())

}