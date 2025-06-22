package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
	"github.com/google/uuid"
)

func followFeed(s *state, cmd command) error {

	userInfo, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedInfo, err := s.db.GetFeedFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	
	createdFeedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID : userInfo.ID,
		FeedID : feedInfo.ID,
	})
	if err != nil {
		return err
	}
	
	// print the name of the feed and the current user
	fmt.Printf("User: %s has now subscribed to the feed: %s\n", createdFeedFollowRow.UsersName, createdFeedFollowRow.FeedName)
	
	return nil
}

func followingFeeds(s *state, cmd command) error {

	userInfo, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	
	feedFollowsForUser, err := s.db.GetFeedFollowsForUser(context.Background(), userInfo.ID)
	if err != nil {
		return err
	}
	fmt.Printf("The feeds for users %s are: \n", userInfo.Name)

	for _, feed:= range feedFollowsForUser {
		fmt.Printf(" * %s", feed.FeedName)
	}
	
	return nil
}