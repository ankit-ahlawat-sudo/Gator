package main

import (
	"context"
	"fmt"

	"github.com/ankit-ahlawat-sudo/Gator/fetch"
	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {

	feed, err := fetch.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err!=nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func addFeed(s *state, cmd command) error {

	user:= s.cfg.CurrentUserName
	userDetails, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("eror fetching the current user")
	}

	if len(cmd.Args) < 2 {
		return fmt.Errorf("the addfeed handler expects two argument, the username and the URL")
	}

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: userDetails.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding the feed to the Database")
	}

	feed, err := fetch.FetchFeed(context.Background(), cmd.Args[1])
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)

	
	return nil
}
