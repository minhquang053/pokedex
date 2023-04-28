package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/minhquang053/pokedex/internal/pokecache"
)

func GetLocations(c *pokecache.Cache, purl *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"

	if purl != nil {
		url = *purl
	}

	if val, ok := c.Get(url); ok {
		locationResp := RespLocationAreas{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocationAreas{}, err
		}
		return locationResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}
	locationAreas := RespLocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return RespLocationAreas{}, err
	}

	c.Add(url, body)

	return locationAreas, nil
}
