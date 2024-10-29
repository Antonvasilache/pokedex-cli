package main

import (
	"os"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandExit(cfg *config, client *pokeapi.Client) error{
	os.Exit(0)
	return nil
}