package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListBerries(pageURL *string) (RespBerries, error) {
	url := baseURL + "/berry"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		berryResp := RespBerries{}
		err := json.Unmarshal(val, &berryResp)
		if err != nil {
			return RespBerries{}, err
		}
		return berryResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespBerries{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespBerries{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespBerries{}, err
	}

	berryResp := RespBerries{}
	err = json.Unmarshal(dat, &berryResp)
	if err != nil {
		return RespBerries{}, err
	}

	c.cache.Add(url, dat)
	return berryResp, nil
}