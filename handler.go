package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.name)
	}
	username := cmd.arguments[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("The username: %s has been set\n", username)
	return nil
}
