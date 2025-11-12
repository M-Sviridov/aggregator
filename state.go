package main

import (
	"github.com/M-Sviridov/aggregator/internal/config"
	"github.com/M-Sviridov/aggregator/internal/database"
)

func newState(cfg *config.Config, db *database.Queries) *state {
	return &state{
		db:  db,
		cfg: cfg,
	}
}
