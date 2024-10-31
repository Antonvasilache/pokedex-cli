package main

import (
	"fmt"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandInspect(cfg *config, client *pokeapi.Client, parameter string) error {
	//check if parameter is empty
	if parameter == "" {
		return fmt.Errorf("you must provide a pokemon name")
	}

	//check if we already have the pokemon
	if _, ok := cfg.Pokedex[parameter]; !ok {		
		fmt.Println() 
		return fmt.Errorf("you have not caught that pokemon")
	}
	
	//print pokemon stats
	fmt.Printf("Name: %s\n", cfg.Pokedex[parameter].Name)
	fmt.Printf("Height: %d\n", cfg.Pokedex[parameter].Height)
	fmt.Printf("Weight: %d\n", cfg.Pokedex[parameter].Weight)
	fmt.Println("Stats:")
	for _, result := range cfg.Pokedex[parameter].Stats {
		fmt.Printf("  - %s: %d\n", result.Stat.Name, result.BaseStat)
	}
	fmt.Println("Types:")
	for _, result := range cfg.Pokedex[parameter].Types {
		fmt.Printf("  - %s\n", result.Type.Name)
	}
	
	return nil
}