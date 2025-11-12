package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/M-Sviridov/aggregator/internal/config"
	"github.com/M-Sviridov/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}
	dbQueries := database.New(db)

	state := newState(&cfg, dbQueries)
	cmds := newCommands()

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(state, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
