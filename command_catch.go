package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon name")
	}
	pokemonName := args[0]

	pokemon, err := c.pokeapiClient.Pokemon(pokemonName)
	if err != nil {
		return err
	}

	if pokemon.Name == "" {
		fmt.Printf("Pokemon %s not found\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	shoot := shootPower(pokemon.BaseExperience)

	if shoot < pokemon.BaseExperience {
		fmt.Printf("%s escaped\n", pokemonName)

		return nil
	}

	c.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}

func shootPower(initialPower int) int {
	funFactor := 30
	min := initialPower - funFactor
	max := initialPower + funFactor
	return rand.Intn(max-min) + min
}
