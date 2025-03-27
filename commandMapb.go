package main

import (
	"fmt"
	"pokedexcli/internal/httpLogic"
	"pokedexcli/internal/jsonTypes"
)

func commandMapb(config *Config) error {
	var locationresp jsonTypes.LocationAreaResponse
	url := config.Previous

	if url == "" {
		fmt.Println("you're on the first page")
		return nil
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
