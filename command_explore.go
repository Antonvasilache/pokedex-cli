package main

import (
	"fmt"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandExplore(cfg *config, client *pokeapi.Client, parameter string) error {
	resp, err := client.GetLocationAreaDetails(parameter)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, result := range resp.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}
	fmt.Println()
	return nil
}