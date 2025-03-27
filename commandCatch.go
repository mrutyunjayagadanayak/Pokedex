package main

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/httpLogic"
	"pokedexcli/internal/jsonTypes"
	"time"
)

func commandCatch(config *Config) error {
	var pokemonResp jsonTypes.PokemonResponse

	if len(config.InputArgs) < 2 {
		fmt.Println("Please provide a pokemon name to catch")
		return fmt.Errorf("No pokemon name provided")
	}

	pokemonName := config.InputArgs[1]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	err := httpLogic.HttpLogic(config.Cache, url, &pokemonResp)
	if err != nil {
		fmt.Println("Incorrect data received")
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

	rand.Seed(time.Now().UnixNano())

	const K = 0.7
	const baseExperienceFactor = 0.001

	probability := K - (float64(pokemonResp.BaseExperience) * baseExperienceFactor)
	catchRoll := rand.Float64()

	if catchRoll < probability {
		fmt.Printf("Caught %s\n", pokemonResp.Name)
		pokedex[pokemonName] = pokemonResp
	} else {
		fmt.Printf("%s ran away\n", pokemonResp.Name)
	}

	return nil
}
