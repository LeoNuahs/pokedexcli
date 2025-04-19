package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListPokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + fmt.Sprintf("/pokemon/%s", pokemonName)

	if data, exists := c.cache.Get(url); exists {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	
	res, err := c.httpClient.Do(req)
	if res.StatusCode >= 300 {
		log.Printf("Response failed with status code: %d\n\n", res.StatusCode)
	}
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
