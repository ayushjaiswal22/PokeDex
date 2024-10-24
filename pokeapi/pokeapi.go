package pokeapi

import (
    "net/http"
    "time"
    "github.com/ayushjaiswal22/pokedexcli/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
    httpClient http.Client
    Next_page_url *string
    Prev_page_url *string
    PokeApiCache pokecache.Cache
}

func NewClient(interval time.Duration) Client {
    return Client {
        httpClient: http.Client{
            Timeout: time.Second * 10,
        },
        Next_page_url:nil,
        Prev_page_url:nil,
        PokeApiCache: pokecache.CreateCache(interval),
    }
}
