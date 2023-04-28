package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/minhquang053/pokedex/internal/pokeapi"
)

func commandCatch(cf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You can only catch one at a time")
	}
	pokemonName := args[0]
	pokemon, err := pokeapi.GetPokemon(cf.Cache, pokemonName)
	if err != nil {
		return err
	}

	if _, ok := cf.Pokedex[pokemon.Name]; !ok {
		cf.Pokedex[pokemon.Name] = pokemon
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

	if res > 40 {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemon.Name)

	return nil
}
