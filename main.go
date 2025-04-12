package main

import (
	"time"

	"github.com/LeoNuahs/pokedexcli/internal/pokeapi"
	"github.com/LeoNuahs/pokedexcli/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(
		5*time.Second,
		pokeCache,
	)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
