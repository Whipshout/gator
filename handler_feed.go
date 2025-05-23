package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/whipshout/gator/internal/database"
)

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get feed owner: %w", err)
		}

		printFeed(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}

func handlerCreateFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed in database: %w", err)
	}

	fmt.Println("Feed created successfully!")
	printFeed(feed, user)

	fmt.Println()
	fmt.Println("=====================================")

	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow in database: %w", err)
	}

	fmt.Println("Feed follow created successfully!")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:				%v\n", feed.ID)
	fmt.Printf(" * Created:			%v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:			%v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:			%v\n", feed.Name)
	fmt.Printf(" * URL:				%v\n", feed.Url)
	fmt.Printf(" * User:			%v\n", user.Name)
	fmt.Printf(" * LastFetchedAt:	%v\n", feed.LastFetchedAt.Time)
}
