package main

import (
	"context"
	"log"

	"github.com/cristiano-pacheco/goflix/cmd"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

// @title           Go modulith API
// @version         0.0.1
// @description     Go modulith API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer <token>**

// @BasePath  /
func main() {
	config.Init()

	cfg := config.GetConfig()
	otel.Init(cfg)

	defer func() {
		if err := otel.Trace().Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	cmd.Execute()
}
