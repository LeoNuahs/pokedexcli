package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n\n")
	os.Exit(0)
	return nil
}