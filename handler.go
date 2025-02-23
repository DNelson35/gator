package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DNelson35/gator/internal/database"
	"github.com/google/uuid"
)


func handlerLogin( s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects a username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	if err := s.pconfig.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("user set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register expects a username")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	})

	if err != nil {
		return err
	}

	if err = s.pconfig.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("user created and loged in: %v", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("this command takes no arguments")
	}
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("this command takes no arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.pconfig.CurrentUserName {
			fmt.Printf("%v (current)\n", user.Name)
			continue
		}
		fmt.Println(user.Name)
	}
	return nil
}

func handleAgg(_ *state, _ command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Println(*rss)
	return nil
}
