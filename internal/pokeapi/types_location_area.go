package pokeapi

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NamedAPIResource      `json:"location"`
	Name                 string                `json:"name"`
	Names                []LocalizedName       `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocalizedName struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource               `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetail `json:"version_details"`
}

type EncounterMethodVersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource                `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type EncounterDetail struct {
	Chance          int              `json:"chance"`
	ConditionValues []interface{}    `json:"condition_values"`
	MaxLevel        int              `json:"max_level"`
	MinLevel        int              `json:"min_level"`
	Method          NamedAPIResource `json:"method"`
}
