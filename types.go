package main

import "github.com/Antonvasilache/pokedex-cli/internal/pokeapi"

type cliCommand struct {
	name		string
	description string
	callback 	func(cfg *config, client *pokeapi.Client) error
}

type config struct {
	Next 		*string
	Previous 	*string
}