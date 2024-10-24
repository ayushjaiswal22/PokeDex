package pokeapi

import (
    "fmt"
    "net/http"
    "io"
    "errors"
    "encoding/json"
)

type area struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListAreaPokemons(areaName string) (area, error) {
    if areaName == "" {
        return area{}, errors.New("area is empty")
    }
    endpoint := "location-area/"
    full_url := baseURL + endpoint + areaName
    req, err := http.NewRequest("GET", full_url, nil)
    if err!=nil {
        return area{}, err
    }
    areaData, ok := c.PokeApiCache.GetCacheVal(full_url)
    if ok {
        fmt.Println("cache hit")
        areaLoc := area{}
        err := json.Unmarshal(areaData, &areaLoc)
        if err!=nil{
            return area{}, err
        }
        c.PokeApiCache.UpdateCacheVal(full_url)
        return areaLoc, nil
    }
    fmt.Println("cache miss")

    resp, err := c.httpClient.Do(req)
    if err!=nil {
        return area{}, err
    }
    defer resp.Body.Close()
    
    data, err := io.ReadAll(resp.Body)
    if err!=nil {
        return area{}, err
    }
    areaLoc := area{}
    err = json.Unmarshal(data, &areaLoc)
    if err!=nil {
        return area{}, err
    }
    c.PokeApiCache.AddCacheVal(full_url, data)
    return areaLoc, nil
}
