package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetMove(moveName string) (Move, error) {
	url := baseURL + "/move/" + moveName
	
	if val, ok := c.cache.Get(url); ok {
		moveResp := Move{}
		err := json.Unmarshal(val, &moveResp)
		if err != nil {
			return Move{}, err
		}
		return moveResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Move{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Move{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Move{}, err
	}

	moveResp := Move{}
	err = json.Unmarshal(dat, &moveResp)
	if err != nil {
		return Move{}, errors.New("error: invalid move")
	}

	c.cache.Add(url, dat)
	return moveResp, nil
}