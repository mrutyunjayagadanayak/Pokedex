package httpLogic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func HttpLogic[T any](cache *pokecache.Cache, url string, target *T) error {

	data, exists := cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("NonOk status code received")
		}

		data, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}
		cache.Add(url, data)
	}

	err := json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
