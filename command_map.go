package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResp, clientErr := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if clientErr != nil {
		return clientErr
	}

	cfg.nextLocationURL = &locationsResp.Next
	cfg.prevLocationURL = &locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	
	fmt.Println()
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil || *cfg.prevLocationURL == "" {
		fmt.Println("You're on the first page")
		
		return nil
	}
	
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}
	
	cfg.nextLocationURL = &locationsResp.Next
	cfg.prevLocationURL = &locationsResp.Previous
	
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	
	fmt.Println()
	return nil
}