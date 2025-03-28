package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/re-jas/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	caughtPokemon    map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all the Pokemon located in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get information about caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get names of all caught pokemons",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
