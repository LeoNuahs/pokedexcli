package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

func main() {
	// List of commands
	commands = map[string]cliCommand{
		"exit": {
			name: "exit", 
			description: "Exit the Pokedex", 
			callback: commandExit,
		},
		"help": {
			name: "help", 
			description: "Displays a help message", 
			callback: commandHelp,
		},
	}

	// Read from system
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			if len(input) == 0 {
				continue
			}
			cleaned := cleanInput(input)

			c, ok := commands[cleaned[0]]
			if !ok {
				fmt.Printf("Unknown command\n\n")
				continue
			}

			err := c.callback()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	lowercased := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowercased)
	cleanedInputs := strings.Fields(trimmed)

	return cleanedInputs
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage: \n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println("")

	return nil
}