package maplogic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func MapLogic[T any](cache *pokecache.Cache, url string, target *T) error {

	data, exists := cache.Get(url)
	if !exists {
		fmt.Println("Date not in cache")
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return err
		}

		data, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}
		cache.Add(url, data)
	} else {
		fmt.Println("Data got from cache")
	}

	err := json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
