package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/jsonTypes"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type Config struct {
	Next      string
	Previous  string
	Cache     *pokecache.Cache
	InputArgs []string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var registry = make(map[string]cliCommand)

var pokedex = make(map[string]jsonTypes.PokemonResponse)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{
		Cache: pokecache.NewCache(10 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		inputSlice := cleanInput(input)
		if len(inputSlice) == 0 {
			continue
		}
		input = inputSlice[0]
		commandValue, exists := registry[input]
		config.InputArgs = inputSlice
		if exists {
			commandValue.callback(config)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	var result []string
	if text == "" {
		return result
	}
	text = strings.ToLower(text)
	result = strings.Fields(text)
	return result
}

func init() {
	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	registry["map"] = cliCommand{
		name:        "map",
		description: "Display a list of areas",
		callback:    commandMap,
	}

	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display the previuous areas",
		callback:    commandMapb,
	}

	registry["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a particular region",
		callback:    commandExplore,
	}

	registry["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a given pokemon",
		callback:    commandCatch,
	}

	registry["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect an already caught Pokemon",
		callback:    commandInspect,
	}

	registry["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Show all caught pokemon",
		callback:    commandPokedex,
	}
}
