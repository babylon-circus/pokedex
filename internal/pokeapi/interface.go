package pokeapi

type APIClient interface {
	Pokemon(name string) (Pokemon, error)
	LocationArea(id string) (LocationArea, error)
	ListLocations(pageURL *string) (RespShallowLocations, error)
}
