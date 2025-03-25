package main

import (
	maplogic "pokedexcli/internal/mapOperations"
)

func commandMap(config *Config) error {
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	locationresp, err := maplogic.MapLogic(config.Cache, url)
	if err != nil {
		return err
	}

	config.Next = locationresp.Next
	config.Previous = locationresp.Previous

	return nil
}
