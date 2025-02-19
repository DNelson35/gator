package main

import (
	"fmt"
	"os"

	config "github.com/DNelson35/gator/internal/config"
)

type state struct {
	pconfig *config.Config
}

func main () {
	var s state
	cfg := config.Read()
	s.pconfig = &cfg

	var c commands
	c.cmds = make(map[string]func(*state, command) error)

	c.register("login",handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	com := command{
		name: args[1],
		args: args[2:],
	}

	err := c.run(&s, com)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	cfg = config.Read()
	fmt.Println(cfg.CurrentUserName)
	os.Exit(0)
}