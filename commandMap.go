package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

func commandMap(config *Config) error {
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	var locationresp LocationAreaResponse
	cache := pokecache.NewCache(10 * time.Second)
	data, exists := cache.Get(url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &locationresp)
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
