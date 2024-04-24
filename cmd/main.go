package main

import (
	"context"
	"os"
	"time"

	"farmers_connect/internal/config"
	"farmers_connect/internal/repo"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    true,
	}).
		With().Caller().Logger().
		Level(zerolog.TraceLevel)

	cfg := &config.Config{}
	if err := config.Read(cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	db, err := repo.NewDB(ctx, cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to db")
	}
	defer db.Close(ctx)
}
