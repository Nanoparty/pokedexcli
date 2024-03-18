package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

func (c *Client) GetBerry(berryName string) (Berry, error) {
	url := baseURL + "/berry/" + berryName
	if val, ok := c.cache.Get(url); ok {
		berryResp := Berry{}
		err := json.Unmarshal(val, &berryResp)
		if err != nil {
			return Berry{}, err
		}
		return berryResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Berry{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Berry{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Berry{}, err
	}

	berryResp := Berry{}
	err = json.Unmarshal(dat, &berryResp)
	if err != nil {
		return Berry{}, err
	}

	c.cache.Add(url, dat)
	return berryResp, nil
}