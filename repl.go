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

func getCommands() map[string]cliCommand {
	// List of commands
	return map[string]cliCommand{
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
}

func startRepl() {
	// Read from system
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Printf("Unknown command\n\n")
			continue
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}