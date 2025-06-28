package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/ankit-ahlawat-sudo/Gator/internal/config"
	"github.com/ankit-ahlawat-sudo/Gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err:= config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}
	
	dbURL := cfg.DbURL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB not opening")
	}

	dbQueries := database.New(db)
	programState.db = dbQueries

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", getUsersHandler)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(addFeed))
	cmds.register("feeds", getFeedsInfo)
	cmds.register("follow", middlewareLoggedIn(followFeed))
	cmds.register("following", middlewareLoggedIn(followingFeeds))
	cmds.register("unfollow", middlewareLoggedIn(deleteFeedFollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowser))

	if len(os.Args) < 2  {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	
	if err:= cmds.run(programState, command{cmdName, cmdArgs}); err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
