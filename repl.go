package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry = make(map[string]cliCommand)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

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
			commandValue.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	text := `Welcome to the Pokedex!
Usage:
	`
	fmt.Println(text)
	for key, val := range registry {
		fmt.Println(key, ":", val.description)
	}
	return nil
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
}
