package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, nil
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
