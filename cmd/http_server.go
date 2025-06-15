package cmd

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"openfinance/configuration/logger"
)

func StartHttpServer() {
	logger.L().Info("Setting up http server.")

	r := chi.NewRouter()

	LoadRoutes(r)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

	logger.L().Info("Starting http server on port 8080")

}
