package main

import (
	"fmt"

	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandPokedex(cfg *config, client *pokeapi.Client, parameter string) error{
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("your Pokedex is empty. Start catching some pokemon")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Println("  - " + pokemon.Name)
	}

	return nil
}