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