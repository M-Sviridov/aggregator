package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.name)
	}
	url := cmd.arguments[0]

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUser)

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't get feed by url: %w\n", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feed.CreatedAt,
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Couldn't get feed follow: %w\n", err)
	}

	for _, entry := range feedFollow {
		fmt.Println(entry.FeedName)
		fmt.Println(s.cfg.CurrentUser)
	}

	return nil
}
