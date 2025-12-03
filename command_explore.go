package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing location name")
	}
	locationName := args[0]

	locationsAreaResp, err := cfg.pokeapiClient.LocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Printf("Found Pokemon: \n")
	for _, pok := range locationsAreaResp.PokemonEncounters {
		fmt.Printf("- %s\n", pok.Pokemon.Name)
	}
	return nil
}
