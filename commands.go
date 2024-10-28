package main

type cliCommand struct {
	name		string
	description string
	callback 	func() error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: callbackExit,
		},
	}
}

		