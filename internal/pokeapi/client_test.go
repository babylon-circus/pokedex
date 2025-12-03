package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/babylon-circus/pokedexcli/internal/pokecache"
)

func TestClient_Pokemon(t *testing.T) {
	tests := []struct {
		name           string
		pokemonName    string
		serverResponse string
		statusCode     int
		wantErr        bool
	}{
		{
			name:        "successful fetch",
			pokemonName: "pikachu",
			serverResponse: `{
				"name": "pikachu",
				"height": 4,
				"weight": 60,
				"base_experience": 112,
				"abilities": [],
				"forms": [],
				"game_indices": [],
				"held_items": [],
				"moves": [],
				"stats": [],
				"types": []
			}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:           "not found",
			pokemonName:    "fakemon",
			serverResponse: `{"error": "not found"}`,
			statusCode:     http.StatusNotFound,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := Client{
				cache:      pokecache.NewCache(time.Minute),
				httpClient: http.Client{Timeout: 5 * time.Second},
				baseURL:    server.URL,
			}

			pokemon, err := client.Pokemon(tt.pokemonName)

			if (err != nil) != tt.wantErr {
				t.Errorf("Pokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && pokemon.Name != tt.pokemonName {
				t.Errorf("Pokemon() name = %v, want %v", pokemon.Name, tt.pokemonName)
			}
		})
	}
}

func TestClient_Pokemon_Cache(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"name": "pikachu",
			"height": 4,
			"weight": 60,
			"base_experience": 112,
			"abilities": [],
			"forms": [],
			"game_indices": [],
			"held_items": [],
			"moves": [],
			"stats": [],
			"types": []
		}`))
	}))
	defer server.Close()

	client := Client{
		cache:      pokecache.NewCache(time.Minute),
		httpClient: http.Client{Timeout: 5 * time.Second},
		baseURL:    server.URL,
	}

	_, err := client.Pokemon("pikachu")
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}

	_, err = client.Pokemon("pikachu")
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected 1 HTTP call (cached), got %d", callCount)
	}
}

func TestClient_LocationArea(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"name": "canalave-city-area",
			"game_index": 1,
			"encounter_method_rates": [],
			"location": {"name": "canalave-city", "url": ""},
			"names": [],
			"pokemon_encounters": []
		}`))
	}))
	defer server.Close()

	client := Client{
		cache:      pokecache.NewCache(time.Minute),
		httpClient: http.Client{Timeout: 5 * time.Second},
		baseURL:    server.URL,
	}

	location, err := client.LocationArea("canalave-city-area")
	if err != nil {
		t.Fatalf("LocationArea() error = %v", err)
	}

	if location.Name != "canalave-city-area" {
		t.Errorf("LocationArea() name = %v, want canalave-city-area", location.Name)
	}
}

func TestClient_ListLocations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"count": 2,
			"next": null,
			"previous": null,
			"results": [
				{"name": "location1", "url": "url1"},
				{"name": "location2", "url": "url2"}
			]
		}`))
	}))
	defer server.Close()

	client := Client{
		cache:      pokecache.NewCache(time.Minute),
		httpClient: http.Client{Timeout: 5 * time.Second},
		baseURL:    server.URL,
	}

	locations, err := client.ListLocations(nil)
	if err != nil {
		t.Fatalf("ListLocations() error = %v", err)
	}

	if locations.Count != 2 {
		t.Errorf("ListLocations() count = %v, want 2", locations.Count)
	}

	if len(locations.Results) != 2 {
		t.Errorf("ListLocations() results length = %v, want 2", len(locations.Results))
	}
}
