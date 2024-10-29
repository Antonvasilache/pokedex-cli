package pokeapi

import "github.com/Antonvasilache/pokedex-cli/internal/pokecache"

type LocationsAreaResp struct {
	Count		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous 	*string		`json:"previous"`
	Results		[]Location	`json:"results"`
}

type Location struct {
	Name	string	`json:"name"`
	URL		string	`json:"url"`
}

type Client struct {
	cache 	*pokecache.Cache
	baseURL string
}
