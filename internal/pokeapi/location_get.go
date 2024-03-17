package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	
	if val, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, errors.New("error: invalid location")
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}