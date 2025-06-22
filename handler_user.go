package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
	"github.com/google/uuid"
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

func registerHandler(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	cxt:= context.Background()

	_, err := s.db.GetUser(cxt, cmd.Args[0])
	if err == nil {
		return fmt.Errorf("the user already exists")
	}

	_, err = s.db.CreateUser(cxt, database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})
	if err != nil {
		return err
	}

	fmt.Println("registration done for " + cmd.Args[0])

	loginHandler(s, cmd)

	return nil
}
