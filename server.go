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
	db := mongo.Init(cfg.DB_URI, cfg.DB_NAME)

	if cfg.APP_ENV == "dev" {
		fmt.Printf("configs: %+v\n", cfg)
	}

	router := httprouter.New()

	// Application
	appHandlers := v1.CreateAppHandlers(db)
	router.GET("/app", v1.AuthMiddleware(cfg.AUTH_HOST)(appHandlers.Get))
	router.POST("/app", v1.AuthMiddleware(cfg.AUTH_HOST)(appHandlers.Post))
	router.PATCH("/app/:id", v1.AuthMiddleware(cfg.AUTH_HOST)(appHandlers.Patch))
	router.POST("/app/:id/track", appHandlers.Track)

	fmt.Println("listening on port", cfg.PORT)
	err := http.ListenAndServe(":"+cfg.PORT, router)
	if err != nil {
		panic(err)
	}
}
