package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ankit-ahlawat-sudo/Gator/fetch"
	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("expects a single argument, the time_between_reqs")
	}

	timeBetweenRequests, err:= time.ParseDuration(cmd.Args[0])
	if err!=nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {
	nextFeedURL, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("no feed to fetch")
		return 
	}

	nextFeed, err := s.db.GetFeedFromURL(context.Background(), nextFeedURL)
	if err != nil {
		fmt.Println("no feed to fetch")
		return 
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return 
	}

	feed, err := fetch.FetchFeed(context.Background(), nextFeedURL)
	if err != nil {
		fmt.Println("no feed to fetch")
		return 
	}

	for _, feedItem := range feed.Channel.Item {
		//fmt.Println(feedItem.Title)
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, feedItem.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			Title: feedItem.Title,
			Url: feedItem.Link,
			Description: sql.NullString{
				String: feedItem.Description,
				Valid: true,
			},
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	//fmt.Println("================================================================================================================")

}
