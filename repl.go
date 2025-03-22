package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     string
	Previous string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var registry = make(map[string]cliCommand)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}
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
}
