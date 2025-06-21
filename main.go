package main

import (
	"log"
	"os"

	"github.com/ankit-ahlawat-sudo/Gator/internal/config"
)
type state struct {
	cfg *config.Config
}

func main() {
	cfg, err:= config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", loginHandler)

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
