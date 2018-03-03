package main

import (
	"fmt"

	"github.com/esayemm/analytics/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	// db := mongo.Init(cfg.DB_URI, cfg.DB_NAME)

	if cfg.APP_ENV == "dev" {
		fmt.Printf("configs: %+v\n", cfg)
	}

	// r, err := api.New(cfg)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("listening on port", cfg.PORT)
	// http.ListenAndServe(":"+cfg.PORT, r)
}
