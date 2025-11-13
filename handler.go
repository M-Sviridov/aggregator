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

	_, err := s.db.GetUser(context.Background(), username)
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

	_, err := s.db.GetUser(context.Background(), username)
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

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments\n", cmd.name)
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error in deleting all users: %w", err)
	}

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

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments\n", cmd.name)
	}

	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w\n", err)
	}

	fmt.Println(rssFeed)
	return nil
}

func handlerFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: %s <feedname> <feedurl>", cmd.name)
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]
	username := s.cfg.CurrentUser
	userDB, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("Could not get user from DB: %w\n", err)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    userDB.ID,
	}

	s.db.CreateFeed(context.Background(), feedParams)
	fmt.Printf("%+v\n", feedParams)
	return nil
}

func handlerShowFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments\n", cmd.name)
	}

	feeds, err := s.db.GetUserFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get user feeds: %w\n", err)
	}
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", feed.UserName)
	}

	return nil
}
