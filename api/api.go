package api

import (
	"github.com/esayemm/analytics/api/app"
	"github.com/esayemm/analytics/config"
	"github.com/esayemm/analytics/middleware"
	"github.com/go-chi/chi"
)

func New(cfg config.Config) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.Auth(cfg.AUTH_HOST))
	r.Mount("/app", app.Router())

	return r, nil
}
