package main

import (
	"fmt"
)

func commandExplore(cfg *config, target []string) error {
	location := target[0]
	fmt.Printf("Exploring %s...", location)
	
	pokemonsResp, clientErr := cfg.pokeapiClient.ListLocationAreasPokemons(location)
	if clientErr != nil {
		return clientErr
	}

	fmt.Printf("\nFound Pokemon:\n")
	for _, encounter := range pokemonsResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
