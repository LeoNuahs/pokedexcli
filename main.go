package main

import (
	"time"

	"github.com/LeoNuahs/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemons: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
