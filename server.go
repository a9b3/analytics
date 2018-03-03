package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/esayemm/analytics/api"
	"github.com/esayemm/analytics/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.WithError(err).Fatalf("config init %s", err.Error())
		return
	}
	// db := mongo.Init(cfg.DB_URI, cfg.DB_NAME)

	if cfg.APP_ENV == "dev" {
		logrus.Infof("configs: %+v\n", cfg)
	}

	r, err := api.New(cfg)
	if err != nil {
		logrus.WithError(err).Fatalf("api init %s", err.Error())
		return
	}

	logrus.Infof("listening on port %s", cfg.PORT)
	if err := http.ListenAndServe(":"+cfg.PORT, r); err != nil {
		logrus.WithError(err).Fatalf("server init %s", err.Error())
		return
	}
}
