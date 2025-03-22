package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *Config) error {
	url := config.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var locationresp LocationAreaResponse
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
