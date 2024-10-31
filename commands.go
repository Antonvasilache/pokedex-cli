package main

func getCommands() map[string]cliCommand{
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Display the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the previous 20 location areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <area_name>",
			description: "Display the pokemon that can be encountered in the given area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback: commandCatch,
		},
	}
}

		