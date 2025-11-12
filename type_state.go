package main

import (
	"github.com/M-Sviridov/aggregator/internal/config"
	"github.com/M-Sviridov/aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
