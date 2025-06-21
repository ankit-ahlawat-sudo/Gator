package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ankit-ahlawat-sudo/Gator/internal/database"
	"github.com/google/uuid"
)

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
