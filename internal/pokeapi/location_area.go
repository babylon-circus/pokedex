package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationArea(id string) (LocationArea, error) {
	url := baseURL + "/location-area/" + id

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
