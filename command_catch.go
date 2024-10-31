package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func commandCatch(cfg *config, client *pokeapi.Client, parameter string) error{	
	//check if we already have the pokemon
	if _, ok := cfg.Pokedex[parameter]; ok {
		fmt.Printf("%s already exists in your Pokedex\n", parameter)
		fmt.Println() 
		return nil
	}

	//get pokemon info
	resp, err := client.GetPokemonInfo(parameter)
	if err != nil {
		return fmt.Errorf("%w. Please use a valid Pokemon name or id", err)
	}		
		
	//attempt to catch
	fmt.Printf("Throwing a ball at %s...\n", parameter)
	chanceToCatch := 100 - (resp.BaseExperience / 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchToken := r.Intn(100)

	if catchToken <= chanceToCatch {
		cfg.Pokedex[parameter] = resp
		fmt.Printf("%s was caught!\n", parameter)
		fmt.Println()
		return nil
	}
		
	fmt.Printf("%s escaped!\n", parameter)
	fmt.Println()		

	return nil
}
