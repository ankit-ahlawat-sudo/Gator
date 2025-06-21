package main

import (
	"context"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	cxt:= context.Background()

	_, err := s.db.GetUser(cxt, cmd.Args[0])

	if err != nil {
		return fmt.Errorf("the user doesn't exist in the Database")
	}


	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println( "User switched successfully to " +  cmd.Args[0])

	return nil
}