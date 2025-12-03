package pokeapi

import (
	"fmt"
)

func (c *Client) LocationArea(id string) (LocationArea, error) {
	url := c.baseURL + "/location-area/" + id

	locationArea, err := fetchAndCache[LocationArea](c, url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to fetch location area %s: %w", id, err)
	}

	return locationArea, nil
}
