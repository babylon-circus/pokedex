package main

import (
	"time"

	"github.com/babylon-circus/pokedexcli/internal/pokeapi"
	"github.com/babylon-circus/pokedexcli/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: &pokeClient,
		pokedex:       pokedex.New(),
	}

	startRepl(cfg)
}
