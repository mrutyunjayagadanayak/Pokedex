package main

import (
	"fmt"
	"pokedexcli/internal/jsonTypes"
	maplogic "pokedexcli/internal/mapOperations"
)

func commandMap(config *Config) error {
	var locationresp jsonTypes.LocationAreaResponse

	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	err := maplogic.MapLogic(config.Cache, url, &locationresp)
	if err != nil {
		return err
	}

	config.Next = locationresp.Next
	config.Previous = locationresp.Previous
	for _, result := range locationresp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
