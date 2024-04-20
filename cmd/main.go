package main

import (
	"context"
	"farmers_connect/internal/config"
	"farmers_connect/internal/repo"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := &config.Config{}
	if err := config.Read(cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	db, err := repo.NewDB(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)
}
