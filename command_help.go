package main

import (
	"fmt"
)

func commandHelp(cf *config, args ...string) error {
	commands := getCommands()
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range commands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	fmt.Printf("\n")
	return nil
}
