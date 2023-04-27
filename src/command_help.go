package main

import (
	"fmt"
)

func commandHelp(cf *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Display a help message\nexit:Exit the pokedex\n\n")
	return nil
}
