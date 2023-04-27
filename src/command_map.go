package main

import (
	"./internal/pokeapi"
)

func commandMapf(cf *config) error {
	getLocations(cf.Next)
}

func commandMapb(cf *config) error {
	getLocations(cf.Previous)
}
