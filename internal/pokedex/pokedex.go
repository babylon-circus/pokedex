package pokedex

import (
	"errors"
	"fmt"
	"sync"

	"github.com/babylon-circus/pokedexcli/internal/pokeapi"
)

var (
	ErrPokemonNotCaught = errors.New("pokemon not caught")
	ErrAlreadyCaught    = errors.New("pokemon already caught")
)

type Pokedex struct {
	pokemon map[string]pokeapi.Pokemon
	mu      sync.RWMutex
}

func New() *Pokedex {
	return &Pokedex{
		pokemon: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Catch(pokemon pokeapi.Pokemon) error {
	if pokemon.Name == "" {
		return errors.New("invalid pokemon: name is empty")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.pokemon[pokemon.Name]; exists {
		return fmt.Errorf("%w: %s", ErrAlreadyCaught, pokemon.Name)
	}

	p.pokemon[pokemon.Name] = pokemon
	return nil
}

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	pokemon, exists := p.pokemon[name]
	if !exists {
		return pokeapi.Pokemon{}, fmt.Errorf("%w: %s", ErrPokemonNotCaught, name)
	}

	return pokemon, nil
}

func (p *Pokedex) Has(name string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	_, exists := p.pokemon[name]
	return exists
}

func (p *Pokedex) List() []pokeapi.Pokemon {
	p.mu.RLock()
	defer p.mu.RUnlock()

	list := make([]pokeapi.Pokemon, 0, len(p.pokemon))
	for _, pokemon := range p.pokemon {
		list = append(list, pokemon)
	}

	return list
}

func (p *Pokedex) Count() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.pokemon)
}
