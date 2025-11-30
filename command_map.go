package main

import (
	"errors"
	"fmt"
)

func commandMapF(c *config) error {

	areaLocation := c.pokeapiClient.GetAreLocation(c.nextLocationsURL)
	c.nextLocationsURL = areaLocation.Next
	c.prevLocationsURL = areaLocation.Previous

	for _, location := range areaLocation.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	areaLocation := c.pokeapiClient.GetAreLocation(c.prevLocationsURL)
	c.nextLocationsURL = areaLocation.Next
	c.prevLocationsURL = areaLocation.Previous

	for _, location := range areaLocation.Results {
		fmt.Println(location.Name)
	}

	return nil
}
