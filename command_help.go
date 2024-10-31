package main

import (
	"fmt"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandHelp(cfg *config, client *pokeapi.Client, parameter string) error{
	fmt.Println()
	fmt.Println("welcome to the Pokedex")
 	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}