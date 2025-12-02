package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
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
