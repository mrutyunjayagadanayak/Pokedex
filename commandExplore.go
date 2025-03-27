package main

import (
	"fmt"
	"pokedexcli/internal/httpLogic"
	"pokedexcli/internal/jsonTypes"
)

func commandExplore(config *Config) error {
	var pokemonAreaResp jsonTypes.PokemonListAreaResponse

	if len(config.InputArgs) < 2 {
		fmt.Println("Please provide a location area name")
		return fmt.Errorf("No area provided to explore")
	}
	areaName := config.InputArgs[1]

	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	err := httpLogic.HttpLogic(config.Cache, url, &pokemonAreaResp)
	if err != nil {
		println("Incorrect data received")
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonAreaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
