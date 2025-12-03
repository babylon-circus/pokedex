package pokedex

import (
	"errors"
	"testing"

	"github.com/babylon-circus/pokedexcli/internal/pokeapi"
)

func TestPokedex_Catch(t *testing.T) {
	tests := []struct {
		name        string
		pokemon     pokeapi.Pokemon
		wantErr     bool
		expectedErr error
	}{
		{
			name: "catch valid pokemon",
			pokemon: pokeapi.Pokemon{
				Name:   "pikachu",
				Height: 4,
				Weight: 60,
			},
			wantErr: false,
		},
		{
			name: "catch pokemon with empty name",
			pokemon: pokeapi.Pokemon{
				Name:   "",
				Height: 4,
			},
			wantErr: true,
		},
		{
			name: "catch already caught pokemon",
			pokemon: pokeapi.Pokemon{
				Name:   "pikachu",
				Height: 4,
			},
			wantErr:     true,
			expectedErr: ErrAlreadyCaught,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New()

			if tt.name == "catch already caught pokemon" {
				_ = p.Catch(tt.pokemon)
			}

			err := p.Catch(tt.pokemon)

			if (err != nil) != tt.wantErr {
				t.Errorf("Catch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Catch() error = %v, expectedErr %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func TestPokedex_Get(t *testing.T) {
	p := New()
	pikachu := pokeapi.Pokemon{
		Name:   "pikachu",
		Height: 4,
		Weight: 60,
	}
	_ = p.Catch(pikachu)

	tests := []struct {
		name        string
		pokemonName string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "get existing pokemon",
			pokemonName: "pikachu",
			wantErr:     false,
		},
		{
			name:        "get non-existent pokemon",
			pokemonName: "charizard",
			wantErr:     true,
			expectedErr: ErrPokemonNotCaught,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Get(tt.pokemonName)

			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.Name != tt.pokemonName {
				t.Errorf("Get() got = %v, want %v", got.Name, tt.pokemonName)
			}

			if tt.wantErr && tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Get() error = %v, expectedErr %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func TestPokedex_Has(t *testing.T) {
	p := New()
	pikachu := pokeapi.Pokemon{Name: "pikachu"}
	_ = p.Catch(pikachu)

	tests := []struct {
		name        string
		pokemonName string
		want        bool
	}{
		{
			name:        "has existing pokemon",
			pokemonName: "pikachu",
			want:        true,
		},
		{
			name:        "does not have pokemon",
			pokemonName: "charizard",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Has(tt.pokemonName); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokedex_Count(t *testing.T) {
	p := New()

	if got := p.Count(); got != 0 {
		t.Errorf("Count() = %v, want 0", got)
	}

	_ = p.Catch(pokeapi.Pokemon{Name: "pikachu"})
	if got := p.Count(); got != 1 {
		t.Errorf("Count() = %v, want 1", got)
	}

	_ = p.Catch(pokeapi.Pokemon{Name: "charizard"})
	if got := p.Count(); got != 2 {
		t.Errorf("Count() = %v, want 2", got)
	}
}

func TestPokedex_List(t *testing.T) {
	p := New()
	pikachu := pokeapi.Pokemon{Name: "pikachu"}
	charizard := pokeapi.Pokemon{Name: "charizard"}

	_ = p.Catch(pikachu)
	_ = p.Catch(charizard)

	list := p.List()

	if len(list) != 2 {
		t.Errorf("List() length = %v, want 2", len(list))
	}

	names := make(map[string]bool)
	for _, pokemon := range list {
		names[pokemon.Name] = true
	}

	if !names["pikachu"] || !names["charizard"] {
		t.Errorf("List() missing expected pokemon")
	}
}

func TestPokedex_Concurrency(t *testing.T) {
	p := New()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			pokemon := pokeapi.Pokemon{Name: string(rune('a' + id))}
			_ = p.Catch(pokemon)
			_ = p.Has(pokemon.Name)
			_, _ = p.Get(pokemon.Name)
			_ = p.List()
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if count := p.Count(); count != 10 {
		t.Errorf("Concurrency test: Count() = %v, want 10", count)
	}
}
