package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't parse time duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}

	fmt.Println("Found a feed to fetch!")

	if err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("coudln't collect feed %s: %w", feed.Name, err)
	}

	for _, item := range data.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	fmt.Printf("Feed %s collected and %v posts found in total", feed.Name, len(data.Channel.Item))
	return nil
}
