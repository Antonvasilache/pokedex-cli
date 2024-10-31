package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/Antonvasilache/pokedex-cli/internal/pokecache"
)

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(cacheInterval),
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c* Client) GetLocationAreas(nextURL *string) (LocationsAreaResp, error) {
	endpoint := fmt.Sprintf("%s/location-area?offset=0&limit=20", c.baseURL)
	if nextURL != nil && *nextURL != "" {
		endpoint = *nextURL
	}

	//check the cache	
	data, ok := c.cache.Get(endpoint)
	if ok {
		//cache hit
		var locations LocationsAreaResp		
		err := json.Unmarshal(data, &locations)
		if err != nil {
		return LocationsAreaResp{}, err
		}
		return locations, nil
	}	

	res, err := http.Get(endpoint)
	if err != nil {
		return LocationsAreaResp{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationsAreaResp{}, err
	}
	var locations LocationsAreaResp		
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsAreaResp{}, err
	}
	
	c.cache.Add(endpoint, data)

	return locations, nil
}

func (c* Client) GetLocationAreaDetails(locationAreaNameOrId string) (LocationAreaDetails, error){
	endpoint := fmt.Sprintf("%s/location-area/%s", c.baseURL, locationAreaNameOrId)

	//check the cache
	data, ok := c.cache.Get(endpoint)
	if ok {
		//cache hit
		//fmt.Println("Cache hit!")
		var details LocationAreaDetails
		err := json.Unmarshal(data, &details)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		return details, nil
	}

	res, err := http.Get(endpoint)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	var details LocationAreaDetails
		err = json.Unmarshal(data, &details)
		if err != nil {
			return LocationAreaDetails{}, err
		}

	c.cache.Add(endpoint, data)

	return details, nil
}

func (c* Client) GetPokemonInfo(pokemonNameOrId string)(Pokemon, error){
	endpoint := fmt.Sprintf("%s/pokemon/%s", c.baseURL, pokemonNameOrId)

	//check the cache
	data, ok := c.cache.Get(endpoint)
	if ok {
		//cache hit
		var info Pokemon
		err := json.Unmarshal(data, &info)
		if err != nil {
			return Pokemon{}, err
		}
		return info, nil
	}

	res, err := http.Get(endpoint)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var info Pokemon
		err = json.Unmarshal(data, &info)
		if err != nil {
			return Pokemon{}, err
		}

	c.cache.Add(endpoint, data)

	return info, nil
}