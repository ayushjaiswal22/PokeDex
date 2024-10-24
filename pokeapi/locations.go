package pokeapi

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
)

type locations struct {
	Count    int         `json:"count"`
	Next     *string      `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(page_url *string) (locations, error) {
    var full_url string
    if page_url == nil {
        endpoint := "/location-area"
        full_url = baseURL + endpoint
    } else {
        full_url = *page_url
    } 
    res, ok := c.PokeApiCache.GetCacheVal(full_url)
    if ok {
        fmt.Println("cache hit")
        locationPoke := locations{}
        err := json.Unmarshal(res, &locationPoke)
        if err!=nil {
            return locations{}, err
        }
        c.Next_page_url = locationPoke.Next
        c.Prev_page_url = locationPoke.Previous
        return locationPoke, nil
    }
    fmt.Println("cache miss")
    req, err := http.NewRequest("GET", full_url, nil)
    if err!=nil {
        return locations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err!=nil {
        return locations{}, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    c.PokeApiCache.AddCacheVal(full_url, data)
    if err!=nil {
        return locations{}, err
    }
    locationPoke := locations{}
    err = json.Unmarshal(data, &locationPoke)
    if err!=nil {
        return locations{}, err
    }
    c.Next_page_url = locationPoke.Next
    c.Prev_page_url = locationPoke.Previous
    return locationPoke, nil
}


