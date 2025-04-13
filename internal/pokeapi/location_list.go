package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResourceList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.cache.Get(url); exists {
		resourceList := ResourceList{}
		err := json.Unmarshal(data, &resourceList)
		if err != nil {
			return ResourceList{}, err
		}

		return resourceList, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResourceList{}, err
	}

	res, err := c.httpClient.Do(req)
	if res.StatusCode >= 300 {
		log.Printf("Response failed with status code: %d\n\n", res.StatusCode)
		return ResourceList{}, err
	}
	if err != nil {
		return ResourceList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResourceList{}, err
	}

	resourceList := ResourceList{}
	err = json.Unmarshal(data, &resourceList)
	if err != nil {
		return ResourceList{}, err
	}

	c.cache.Add(url, data)

	return resourceList, nil
}
