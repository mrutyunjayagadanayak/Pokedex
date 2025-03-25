package main

import (
	"fmt"
	maplogic "pokedexcli/internal/mapOperations"
)

func commandMapb(config *Config) error {
	url := config.Previous

	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationresp, err := maplogic.MapLogic(config.Cache, url)
	if err != nil {
		return err
	}

	config.Next = locationresp.Next
	config.Previous = locationresp.Previous

	return nil
}
