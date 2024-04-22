package main

import (
	"context"
	"github.com/rs/zerolog"
	"log"

	"farmers_connect/internal/config"
	"farmers_connect/internal/repo"
)

func main() {
	ctx := context.Background()

	logger := zerolog.Ctx(ctx)

	cfg := &config.Config{}
	if err := config.Read(cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	db, err := repo.NewDB(ctx, cfg.DB)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create db")
	}
	defer db.Close(ctx)
}
