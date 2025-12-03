package pokeapi

import (
	"fmt"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locations, err := fetchAndCache[RespShallowLocations](c, url)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to list locations: %w", err)
	}

	return locations, nil
}
