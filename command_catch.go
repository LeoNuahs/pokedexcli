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
	pokemon, clientErr := cfg.pokeapiClient.ListPokemonDetails(name)
	if clientErr != nil {
		return clientErr
	}
	
	score := rand.IntN(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if score > 40 {
		fmt.Printf("%s escaped!\n\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n\n", pokemon.Name)
	cfg.caughtPokemons[pokemon.Name] = pokemon

	return nil
}
