package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments", cmd.name)
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error in deleting all users: %w", err)
	}

	return nil
}
