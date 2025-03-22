package main

import "fmt"

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
