package main

import (
	"errors"
	"fmt"
	"github.com/minhquang053/pokedex/internal/pokeapi"
)

func commandMapf(cf *config) error {
	locationResp, err := pokeapi.GetLocations(cf.Next)

	cf.Next = locationResp.Next
	cf.Previous = locationResp.Previous

	if err != nil {
		return err
	}

	for _, res := range locationResp.Results {
		fmt.Println(res.Name)
	}
	return nil
}

func commandMapb(cf *config) error {
	if cf.Previous == nil {
		return errors.New("This is the first page")
	}

	locationResp, err := pokeapi.GetLocations(cf.Previous)

	cf.Next = locationResp.Next
	cf.Previous = locationResp.Previous

	if err != nil {
		return err
	}

	for _, res := range locationResp.Results {
		fmt.Println(res.Name)
	}

	return nil
}
