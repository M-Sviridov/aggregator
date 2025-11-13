package main

import "errors"

func (c *commands) run(s *state, cmd command) error {
	if s.cfg == nil {
		return errors.New("State does not exist")
	}

	handler, ok := c.cmd[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd[name] = f
}

func newCommands() *commands {
	c := &commands{
		cmd: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerFeed)
	return c
}
