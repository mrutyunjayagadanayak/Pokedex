package maplogic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/jsonTypes"
	"pokedexcli/internal/pokecache"
)

func MapLogic(cache *pokecache.Cache, url string) (jsonTypes.LocationAreaResponse, error) {
	var locationresp jsonTypes.LocationAreaResponse
	data, exists := cache.Get(url)
	if !exists {
		fmt.Println("Date not in cache")
		res, err := http.Get(url)
		if err != nil {
			return jsonTypes.LocationAreaResponse{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return jsonTypes.LocationAreaResponse{}, err
		}

		data, err = io.ReadAll(res.Body)

		if err != nil {
			return jsonTypes.LocationAreaResponse{}, err
		}
		cache.Add(url, data)
	} else {
		fmt.Println("Data got from cache")
	}

	err := json.Unmarshal(data, &locationresp)
	if err != nil {
		return jsonTypes.LocationAreaResponse{}, err
	}

	for _, result := range locationresp.Results {
		fmt.Println(result.Name)
	}
	return locationresp, nil
}
