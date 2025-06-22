package fetch

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(cxt context.Context, feedUrl string) (*RSSFeed, error){

	httpReq, err:= http.NewRequestWithContext(cxt, http.MethodGet, feedUrl, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("User-Agent", "gator")
	client:= http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rssFeed RSSFeed

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err:= xml.Unmarshal(body, &rssFeed); err!=nil {
		return nil, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}