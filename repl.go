package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nanoparty/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pokedex")
	fmt.Println("---------------")

	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand {
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"map": {
			name:			"map",
			description:	"Display list of locations",
			callback:		commandMapf,
		},
		"mapb": {
			name:			"mapb",
			description: 	"Display previous list of locations",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Explore a location",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Attempt to catch a Pokemon",
			callback:		commandCatch,
		},
	}
}