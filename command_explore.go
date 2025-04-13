package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	fmt.Printf("Exploring %s...", name)
	
	pokemonsResp, clientErr := cfg.pokeapiClient.ListLocationAreasPokemons(name)
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
