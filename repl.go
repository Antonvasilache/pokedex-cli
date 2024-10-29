package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/Antonvasilache/pokedex-cli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin);
	cfg := &config{}
	client := pokeapi.NewClient(time.Hour)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, exists := getCommands()[commandName]
			if exists {
				err := command.callback(cfg, client)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}		
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}