package main

import (
	"github.com/M-Sviridov/aggregator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	state := newState(&cfg)
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
