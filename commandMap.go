package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	var locationresp LocationAreaResponse
	cache := config.Cache
	data, exists := cache.Get(url)

	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}

	err := json.Unmarshal(data, &locationresp)
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
