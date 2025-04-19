package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage: \n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}