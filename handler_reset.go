package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	
	cxt:= context.Background()
	
	err := s.db.DeleteUsers(cxt)

	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}