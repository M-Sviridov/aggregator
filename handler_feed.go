package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: %s <feedname> <feedurl>", cmd.name)
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]
	username := s.cfg.CurrentUser
	userDB, err := s.db.GetUserByName(context.Background(), username)
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

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feedParams.CreatedAt,
		UpdatedAt: time.Now(),
		UserID:    userDB.ID,
		FeedID:    feedParams.ID,
	}

	s.db.CreateFeed(context.Background(), feedParams)
	s.db.CreateFeedFollow(context.Background(), feedFollowParams)
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
