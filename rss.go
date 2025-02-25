package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/DNelson35/gator/internal/database"
	"github.com/google/uuid"
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

func handlerAddFeed(s *state, cmd command, user database.User) error{
	if len(cmd.args) < 2 {
		return fmt.Errorf("please provide a name and url")
	}
	// user, err := s.db.GetUser(context.Background(), s.pconfig.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}
	printFeed(feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("command takes no arguments")
	}
	feeds, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("command takes one argument: url")
	}
	// user , err := s.db.GetUser(context.Background(), s.pconfig.CurrentUserName)
	// if err != nil {
	// 	return err
	// }
	feed , err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	followedFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	println(followedFeed.FeedName)
	println(followedFeed.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("command takes no arguments")
	}
	// user, err := s.db.GetUser(context.Background(), s.pconfig.CurrentUserName)
	// if err != nil {
	// 	return err
	// }
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range followedFeeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}

func printFeed(feed database.Feed){
	fmt.Println(feed.ID)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
}