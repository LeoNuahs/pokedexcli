package main

import (
	"time"

	"github.com/LeoNuahs/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	
	startRepl(cfg)
}
