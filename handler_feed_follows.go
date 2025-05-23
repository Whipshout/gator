package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/whipshout/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed url: %w", err)
	}

	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	ff, err := s.db.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow in database: %w", err)
	}

	fmt.Println("Feed followed successfully!")
	printFeedFollow(ff, user)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	unfollowParams := database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    url,
	}

	err := s.db.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}

	fmt.Println("Feed unfollowed successfully!")

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ffs, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows for user: %w", err)
	}

	if len(ffs) == 0 {
		fmt.Println("User doesn't follow any feed")
		return nil
	}

	for _, ff := range ffs {
		fmt.Println(ff.FeedName)
	}

	return nil
}

func printFeedFollow(ff database.CreateFeedFollowRow, user database.User) {
	fmt.Printf(" * Feed name:			%v\n", ff.FeedName)
	fmt.Printf(" * Current user:		%v\n", user.Name)
}
