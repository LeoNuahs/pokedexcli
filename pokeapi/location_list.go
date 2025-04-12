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
	
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return RespShallowLocation{}, reqErr
	}
	
	res, resErr := c.httpClient.Do(req)
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s", res.StatusCode, res.Body)
	}
	if resErr != nil {
		log.Fatal(resErr)
		return RespShallowLocation{}, resErr
	}
	
	defer res.Body.Close()

	data, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return RespShallowLocation{}, readErr
	}
	
	locationsResp := RespShallowLocation{}
	decErr := json.Unmarshal(data, &locationsResp)
	if decErr != nil {
		log.Fatal(decErr)
		return RespShallowLocation{}, decErr
	}
	
	// for _, locationArea := range locationAreas.Results {
	// 	fmt.Printf("%v\n", locationArea.Name)
	// }
	
	return locationsResp, nil
}
