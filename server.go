package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/esayemm/analytics/api"
	"github.com/esayemm/analytics/config"
	"github.com/esayemm/analytics/db"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.WithError(err).Fatalf("config init %s", err.Error())
		return
	}
	if cfg.APP_ENV == "dev" {
		logrus.Infof("configs: %+v\n", cfg)
	}

	mdb, err := db.Init(cfg.MONGO_URI, cfg.MONGO_DB_NAME)
	if err != nil {
		logrus.WithError(err).Fatalf("db init %s", err.Error())
		return
	}

	r, err := api.New(&api.APIOption{Mdb: mdb, Cfg: cfg})
	if err != nil {
		logrus.WithError(err).Fatalf("api init %s", err.Error())
		return
	}

	logrus.Infof("listening on port %s", cfg.PORT)
	err = http.ListenAndServe(":"+cfg.PORT, r)
	if err != nil {
		logrus.WithError(err).Fatalf("server init %s", err.Error())
		return
	}
}
