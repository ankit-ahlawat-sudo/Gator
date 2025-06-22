package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, cmd command) error {

	user:= s.cfg.CurrentUserName
	userDetails, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("eror fetching the current user")
	}

	if len(cmd.Args) < 2 {
		return fmt.Errorf("the addfeed handler expects two argument, the username and the URL")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: userDetails.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding the feed to the Database")
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID : userDetails.ID,
		FeedID : feed.ID,
	})
	if err != nil {
		return err
	}


	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	
	return nil
}


func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func getFeedsInfo(s *state, cmd command) error {
	
	getFeedsFromDbRow, err:= s.db.GetFeedsFromDb(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range getFeedsFromDbRow {
		fmt.Printf("feed Name: %s, the URL: %s, created by: %s \n", feed.FeedName, feed.Url, feed.UserName)
		fmt.Println("=====================================")
	}
	
	return nil
}