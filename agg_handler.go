package main

import (
	"context"
	"fmt"
	"github.com/ankit-ahlawat-sudo/Gator/fetch"
)

func aggregator(s *state, cmd command) error {

	feed, err := fetch.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err!=nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
