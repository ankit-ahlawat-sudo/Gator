package main

import (
	"context"
	"fmt"
)

func getUsersHandler(s *state, cmd command) error {
	
	cxt:= context.Background()
	users, err := s.db.GetUsers(cxt)
	if err != nil {
		return err
	}

	currentUser:= s.cfg.CurrentUserName

	for _, user := range users {
		fmt.Print("* " + user.Name)
		if user.Name == currentUser {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}