package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonResp, clientErr := cfg.pokeapiClient.ListPokemonDetails(name)
	if clientErr != nil {
		return clientErr
	}

	// Logic for catching pokemon

	if rand.Int() > pokemonResp.BaseExperience {
		fmt.Printf("%s was caught!\n\n", pokemonResp.Name)
		cfg.caughtPokemons[pokemonResp.Name] = pokemonResp
		} else {
		fmt.Printf("%s escaped!\n\n", pokemonResp.Name)
	}

	return nil
}
