package main

import (
	"fmt"
)


func handlerLogin( s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	if err := s.pconfig.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("user set")
	return nil
}