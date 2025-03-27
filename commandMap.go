package main

import (
	"fmt"
	"pokedexcli/internal/httpLogic"
	"pokedexcli/internal/jsonTypes"
)

func commandMap(config *Config) error {
	var locationresp jsonTypes.LocationAreaResponse

	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	err := httpLogic.HttpLogic(config.Cache, url, &locationresp)
	if err != nil {
		println("Incorrect data received")
		return err
	}

	config.Next = locationresp.Next
	config.Previous = locationresp.Previous
	for _, result := range locationresp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
