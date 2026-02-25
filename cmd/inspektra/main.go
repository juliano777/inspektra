package main

import (
	"fmt"
	"log"

	"github.com/juliano777/inspektra/internal/cli"
	"github.com/juliano777/inspektra/internal/config"
)

func main() {
	// Parse CLI flags
	flags := cli.ParseFlags()

	// Load configuration file
	cfg, err := config.Load(flags.ConfigPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Temporary output to confirm everything is working
	fmt.Println("Configuration loades successfuly")
	fmt.Printf("Postgres host: %s\n", cfg.Postgres.Host)
	fmt.Printf("Application name: %s\n", cfg.Postgres.ApplicationName)
}
