package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
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

	fmt.Println("Feed follow has been created:")
	fmt.Println(feedFollow.FeedName)
	fmt.Println(feedFollow.UserName)

	return nil
}

func handlerShowFeedsFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not take arguments", cmd.name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get current user feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows were found for current user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}

	return nil
}
