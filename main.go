package main

import (
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
	cmds.register("addfeed", addFeed)

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
