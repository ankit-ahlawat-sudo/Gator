package main

import (
	"context"
	"fmt"
	"os"
)

func resetHandler(s *state, cmd command) error {
	
	cxt:= context.Background()
	
	err := s.db.TruncateUsers(cxt)

	if err != nil {
        fmt.Println("Failed to reset users table:", err)
        os.Exit(1)
    }
    fmt.Println("Users table reset successfully.")
    os.Exit(0)
	
	return nil
}