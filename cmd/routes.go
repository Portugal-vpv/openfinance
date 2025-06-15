package cmd

import (
	"github.com/go-chi/chi/v5"
	"openfinance/internal/handler/consent"
)

func LoadRoutes(r *chi.Mux) {
	r.Route("/openfinance/v1", func(r chi.Router) {
		r.Post("/consents", consent.PostConsent)
	})
}
