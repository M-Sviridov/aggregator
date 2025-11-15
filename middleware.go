package main

import (
	"context"
	"fmt"

	"github.com/M-Sviridov/aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return fmt.Errorf("couldn't get user from DB: %w", err)
		}
		return handler(s, cmd, user)
	}
}
