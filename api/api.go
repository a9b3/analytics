package api

import (
	"github.com/esayemm/analytics/api/app"
	"github.com/esayemm/analytics/config"
	"github.com/esayemm/analytics/database"
	"github.com/esayemm/analytics/middleware"
	"github.com/go-chi/chi"
)

type APIOption struct {
	UserStore        *database.UserStore
	ApplicationStore *database.ApplicationStore
	Cfg              *config.Config
}

func New(apiOption *APIOption) (*chi.Mux, error) {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Auth(apiOption.Cfg.AUTH_HOST))

	// app routes
	r.Mount("/app", app.Router())

	return r, nil
}
