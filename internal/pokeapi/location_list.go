package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return RespShallowLocation{}, err
		}
		
		res, err := c.httpClient.Do(req)
		if res.StatusCode >= 300 {
			log.Fatalf("Response failed with status code: %d and \nbody: %s", res.StatusCode, res.Body)
		}
		if err != nil {
			return RespShallowLocation{}, err
		}
		
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocation{}, err
		}
	}
		
	locationsResp := RespShallowLocation{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	c.cache.Add(url, data)
	
	return locationsResp, nil
}
