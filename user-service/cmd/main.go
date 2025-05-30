package main

import (
	"context"
	"log"

	config "github.com/aftosmiros/moodu/user-service/config"
	"github.com/aftosmiros/moodu/user-service/internal/app"
)

func main() {
	ctx := context.Background()

	cfg := config.New()

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("❌ failed to init app: %v", err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatalf("❌ app failed: %v", err)
	}
}
