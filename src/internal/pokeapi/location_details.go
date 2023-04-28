package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/minhquang053/pokedex/internal/pokecache"
)

func GetLocationDetails(c *pokecache.Cache, locationName string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.Get(url); ok {
		locationResp := LocationDetails{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationDetails{}, err
		}
		return locationResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}
	locationDetail := LocationDetails{}
	err = json.Unmarshal(body, &locationDetail)
	if err != nil {
		return LocationDetails{}, err
	}

	c.Add(url, body)

	return locationDetail, nil
}
