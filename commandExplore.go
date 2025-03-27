package main

import (
	"fmt"
	"pokedexcli/internal/jsonTypes"
	maplogic "pokedexcli/internal/mapOperations"
)

func commandExplore(config *Config) error {
	var pokemonResp jsonTypes.PokemonListAreaResponse

	if len(config.InputArgs) < 2 {
		fmt.Println("Please provide a location area name")
		return fmt.Errorf("No area provided to explore")
	}
	areaName := config.InputArgs[1]

	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	err := maplogic.MapLogic(config.Cache, url, &pokemonResp)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
