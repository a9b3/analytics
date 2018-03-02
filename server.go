package main

import (
	"fmt"
	"net/http"

	"github.com/esayemm/analytics/config"
	"github.com/esayemm/analytics/mongo"
	"github.com/esayemm/analytics/v1"
	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg := config.New()
	mongo.Init(cfg.DB_URI, cfg.DB_NAME)

	if cfg.APP_ENV == "dev" {
		fmt.Printf("configs: %+v\n", cfg)
	}

	router := httprouter.New()
	router.GET("/", v1.HelloHandler)

	fmt.Println("listening on port", cfg.PORT)
	err := http.ListenAndServe(":"+cfg.PORT, router)
	if err != nil {
		panic(err)
	}
}
