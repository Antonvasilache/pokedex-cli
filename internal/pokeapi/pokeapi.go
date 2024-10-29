package pokeapi

import (
	"encoding/json"
	"net/http"
)

func NewClient() *Client {
	return &Client{
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c* Client) GetLocationAreas(nextURL *string) (LocationsAreaResp, error) {
	endpoint := c.baseURL + "/location-area"
	if nextURL != nil && *nextURL != "" {
		endpoint = *nextURL
	}

	res, err := http.Get(endpoint)
	if err != nil {
		return LocationsAreaResp{}, err
	}
	defer res.Body.Close()

	var locations LocationsAreaResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return LocationsAreaResp{}, err
	}

	return locations, nil
}