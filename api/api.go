package api

import (
	"github.com/esayemm/analytics/api/app"
	"github.com/esayemm/analytics/config"
	"github.com/esayemm/analytics/db"
	"github.com/esayemm/analytics/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	mgo "gopkg.in/mgo.v2"
)

type APIOption struct {
	Mdb *mgo.Database
	Cfg *config.Config
}

func New(opt *APIOption) (*chi.Mux, error) {
	r := chi.NewRouter()

	// middlewares
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.Auth(opt.Cfg.AUTH_HOST, opt.Cfg.JWT_SECRET))
	r.Use(middleware.LazyCreateUser(opt.Mdb.C(db.UserColName)))

	// app routes
	r.Mount("/app", app.Router(opt.Mdb.C(db.ApplicationColName)))

	return r, nil
}
