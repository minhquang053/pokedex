package main

import "fmt"

func commandPokedex(cf *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cf.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
