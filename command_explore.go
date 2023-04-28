package main

import (
	"errors"
	"fmt"

	"github.com/minhquang053/pokedex/internal/pokeapi"
)

func commandExplore(cf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	locationName := args[0]
	locationDetail, err := pokeapi.GetLocationDetails(cf.Cache, locationName)
	if err != nil {
		return errors.New("This location does not exist")
	}
	pokemonEncounters := locationDetail.PokemonEncounters

	fmt.Printf("Exploring %v...\nFound Pokemon:\n", locationName)
	for _, val := range pokemonEncounters {
		fmt.Printf(" - %v\n", val.Pokemon.Name)
	}
	return nil
}
