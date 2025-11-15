package main

import (
	"context"
	"fmt"
	"time"

	"github.com/M-Sviridov/aggregator/internal/database"
	"github.com/google/uuid"
)

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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.name)
	}
	url := cmd.arguments[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)

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
