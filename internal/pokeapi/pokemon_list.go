package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreasPokemons(locArea string) (LocationArea, error) {
	url := baseURL + fmt.Sprintf("/location-area/%s", locArea)

	if data, exists := c.cache.Get(url); exists {
		locAreasPokemons := LocationArea{}
		err := json.Unmarshal(data, &locAreasPokemons)
		if err != nil {
			return LocationArea{}, err
		}
		return locAreasPokemons, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if res.StatusCode >= 300 {
		log.Printf("Response failed with status code: %d\n\n", res.StatusCode)
	}
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locAreasPokemons := LocationArea{}
	err = json.Unmarshal(data, &locAreasPokemons)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	
	return locAreasPokemons, nil
}