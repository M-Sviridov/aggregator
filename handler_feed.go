package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: %s <feedname> <feedurl>", cmd.name)
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("couldn't get user from DB: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	fmt.Printf("%+v\n", feedFollow)
	return nil
}

func handlerShowFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments", cmd.name)
	}

	feeds, err := s.db.GetUserFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get user feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds were found.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", feed.UserName)
	}

	return nil
}
