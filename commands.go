package main

import "errors"

type command struct{
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.registeredCommands[cmd.Name]; ok {
		return f(s, cmd)
	}
	return errors.New("command not found") 
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCommands[name] = f
}

