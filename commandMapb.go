package main

import (
	"fmt"
	"pokedexcli/internal/jsonTypes"
	maplogic "pokedexcli/internal/mapOperations"
)

func commandMapb(config *Config) error {
	var locationresp jsonTypes.LocationAreaResponse
	url := config.Previous

	if url == "" {
		fmt.Println("you're on the first page")
		return nil
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
