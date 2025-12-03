package pokeapi

import (
	"fmt"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	url := c.baseURL + "/pokemon/" + name

	pokemon, err := fetchAndCache[Pokemon](c, url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to fetch pokemon %s: %w", name, err)
	}

	return pokemon, nil
}
