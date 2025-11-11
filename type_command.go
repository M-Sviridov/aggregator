package main

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}
