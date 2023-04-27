package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getLocations(purl *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"
	if purl != nil {
		url = *purl
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
	return locationAreas, nil
}
