package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "    "
	fmt.Println(strings.Split(text, " "))
	fmt.Println(len(strings.Split(text, " ")))
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
