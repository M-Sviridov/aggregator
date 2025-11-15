package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.name)
	}
	username := cmd.arguments[0]

	_, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User %s does not exist\n", username)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("The username: %s has been set\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.name)
	}
	username := cmd.arguments[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.db.GetUserByName(context.Background(), username)
	if err == nil {
		return fmt.Errorf("The user %s already exists\n", username)
	}

	_, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("Couldn't create user: %w\n", err)
	}
	s.cfg.SetUser(username)
	fmt.Printf("The user %s has been created\n", username)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments\n", cmd.name)
	}

	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error in getting all users: %w", err)
	}

	for _, user := range users {
		if s.cfg.CurrentUser == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments\n", cmd.name)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("Couldn't get current user: %w\n", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get current user feed follows: %w\n", err)
	}

	for _, f := range feedFollows {
		fmt.Println(f.FeedName)
		fmt.Println(s.cfg.CurrentUser)
	}

	return nil
}
