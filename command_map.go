package main

import (
	"fmt"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandMap(cfg *config, client *pokeapi.Client) error{
	resp, err := client.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println()
	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()
	return nil
}

func commandMapb(cfg *config, client *pokeapi.Client) error{
	if cfg.Previous == nil || *cfg.Previous == "" {
		return fmt.Errorf("no previous pages to display")
	}
	resp, err := client.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println()
	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()
	return nil
}