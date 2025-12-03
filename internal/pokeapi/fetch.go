package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchAndCache[T any](c *Client, url string) (T, error) {
	var result T

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &result); err != nil {
			return result, fmt.Errorf("failed to unmarshal cached data for %s: %w", url, err)
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, fmt.Errorf("failed to create request for %s: %w", url, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result, fmt.Errorf("failed to execute request for %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code %d for %s", resp.StatusCode, url)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response body for %s: %w", url, err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal response for %s: %w", url, err)
	}

	c.cache.Add(url, data)
	return result, nil
}
