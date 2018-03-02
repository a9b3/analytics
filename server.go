package main

import (
	"fmt"

	"github.com/esayemm/analytics/config"
)

func main() {
	cfg := config.New()

	if cfg.APP_ENV == "dev" {
		fmt.Printf("configs: %+v\n", cfg)
	}
}
