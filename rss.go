package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"fmt"
)

type RSSFeed struct {
	Channel struct {
		Title 			string 		`xml:"title"`
		Link 				string 		`xml:"link"`
		Description string 		`xml:"description"`
		Items			  []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title  			string 	`xml:"title"`
	Link   			string 	`xml:"link"`
	Description string 	`xml:"description"`
	PubDate     string 	`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error){
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil ,err
	}
	req.Header.Set("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch feed, status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSSFeed

	if err = xml.Unmarshal(body, &rss); err != nil {
		return nil, err
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for _, item := range rss.Channel.Items {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}


	return &rss, nil
}