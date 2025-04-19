package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	name := args[0]

	pokemon, exists := cfg.caughtPokemons[name]
	if !exists {
		return fmt.Errorf("you have not caught a pokemon named %s yet", name)
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %v\n", pokeType.Type.Name)
	}
	fmt.Println("")

	return nil
}
