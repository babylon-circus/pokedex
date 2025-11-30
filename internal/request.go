package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokeLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) get(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)

	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func (c Client) GetAreLocation(url string) PokeLocationArea {
	startUrl := baseURL + "/location-area"
	if url == nil {
		url = startUrl
	}

	result := PokeLocationArea{}
	data := c.get(*url)

	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
