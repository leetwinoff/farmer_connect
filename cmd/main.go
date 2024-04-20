package main

import (
	//Standard library packages
	"context"
	"log"

	//Internal packages
	"farmers_connect/internal/config"
	"farmers_connect/internal/repo"
	//External packages
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
