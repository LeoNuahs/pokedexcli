package main

import (
	"fmt"
)

func commandMapf(cfg *config, target []string) error {
	locationsResp, clientErr := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
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

func commandMapb(cfg *config, target []string) error {
	if cfg.prevLocationURL == nil || *cfg.prevLocationURL == "" {
		fmt.Printf("You're on the first page\n\n")
		return nil
	}

	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
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
