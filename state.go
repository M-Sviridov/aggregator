package main

import "github.com/M-Sviridov/aggregator/internal/config"

func newState(cfg *config.Config) *state {
	return &state{
		cfg: cfg,
	}
}
