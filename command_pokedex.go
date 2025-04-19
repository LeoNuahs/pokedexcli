package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// if len(cfg.caughtPokemons) == 0 {
	// 	fmt.Printf("No caught Pokemons\n\n")
	// 	return nil
	// }

	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", name)
	}
	fmt.Println()

	return nil
}
