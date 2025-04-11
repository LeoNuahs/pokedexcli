package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type LocationAreaResponse struct {
	// Count    int            `json:"count"`
	// Next     string         `json:"next"`
	// Previous string         `json:"previous"`
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	// Url  string `json:"url"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != "" {
		url = cfg.Next
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s", res.StatusCode, res.Body)
	}
	if resErr != nil {
		log.Fatal(resErr)
		return resErr
	}

	data, readErr := io.ReadAll(res.Body)
	defer res.Body.Close()

	if readErr != nil {
		log.Fatal(readErr)
		return readErr
	}

	cfgErr := json.Unmarshal(data, &cfg)
	if cfgErr != nil {
		log.Fatal(cfgErr)
		return cfgErr
	}

	var locationAreas LocationAreaResponse
	decErr := json.Unmarshal(data, &locationAreas)
	if decErr != nil {
		log.Fatal(decErr)
		return decErr
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Printf("%v\n", locationArea.Name)
	}

	return nil
}
