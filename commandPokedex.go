package main

import "fmt"

func commandPokedex(config *Config) error {
	if len(pokedex) == 0 {
		fmt.Println("Till now you have caught no pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range pokedex {
		fmt.Printf("  -%s\n", key)
	}
	return nil
}
