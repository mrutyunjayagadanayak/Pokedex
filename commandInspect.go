package main

import "fmt"

func commandInspect(config *Config) error {
	if len(config.InputArgs) < 2 {
		fmt.Println("Please provide a pokemon name to catch")
		return fmt.Errorf("No pokemon name provided")
	}

	pokemonName := config.InputArgs[1]

	pokemonData, exists := pokedex[pokemonName]

	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name:%s\n", pokemonData.Name)
	fmt.Printf("Height:%d\n", pokemonData.Height)
	fmt.Printf("Weight:%d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemonData.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
