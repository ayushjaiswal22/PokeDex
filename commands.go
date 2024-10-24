package main

import (
    "fmt"
    "time"
    "os"
    "log"
    "github.com/ayushjaiswal22/pokedexcli/pokeapi"
    "math/rand"
)

type cliCommand struct {
    name        string
    description string
    callback    func(string)
}

type Pokedex struct {
    pokedexData map[string]pokeInfo
}

type pokeInfo struct{
    pokemonData pokeapi.Pokemon
    caughtAt time.Time
}

var commandList map[string]cliCommand 
var pokeApiClient pokeapi.Client
var pokeDex Pokedex

func init_func() {
   pokeApiClient = pokeapi.NewClient(time.Second * 10)
   pokeDex = Pokedex{ pokedexData: make(map[string]pokeInfo) }
}
func commandHelp(input string) {
    fmt.Println("\nWelcome to Pokedex!")
    fmt.Println("Usage:\n")
    for cmd, cmdInfo := range commandList {
        fmt.Printf("Command : %s\nDescription: %s\n\n", cmd, cmdInfo.description);
    } 
}

func commandExit(input string){
    fmt.Println("Bye");
    os.Exit(0);
}

func commandMap(input string){
    if pokeApiClient.Prev_page_url != nil && pokeApiClient.Next_page_url == nil {
        fmt.Println("On the last page")
        return 
    }
    resp, err := pokeApiClient.ListLocationAreas(pokeApiClient.Next_page_url)
    if err!=nil {
        log.Fatal(err)
    }

    fmt.Println("Location Areas:")
    for _, area := range resp.Results{
        fmt.Printf(" - %s\n", area.Name)
    }

}

func commandMapBack(input string) {
    if pokeApiClient.Prev_page_url == nil {
        fmt.Println("Already on the first page")
        return 
    }
    resp, err := pokeApiClient.ListLocationAreas(pokeApiClient.Prev_page_url)
    if err!=nil {
        log.Fatal(err)
    }
    
    fmt.Println("Location Areas:")
    for _, area := range resp.Results{
        fmt.Printf(" - %s\n", area.Name)
    }

}

func commandExplore(area string) {
    resp, err := pokeApiClient.ListAreaPokemons(area)
    if err!=nil {
        log.Fatal(err)
    }

    fmt.Printf("Pokemons in %s:\n", area)
    for _, pokemon := range resp.PokemonEncounters {
        fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
    }
}

func commandCatch(pokemon string) {
    if pokemon == "" {
        return 
    }
    _, ok := pokeDex.pokedexData[pokemon]
    if ok {
        fmt.Printf("%s already caught\n", pokemon)
        return
    }
    resp, err := pokeApiClient.GetPokemonInfo(pokemon)
    if err!=nil{
        log.Fatal(err)
    }
    baseExp := resp.BaseExperience
    
    power := rand.Intn(baseExp)
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
    fmt.Printf("Pokemon XP:%d\nTrainer XP:%d\nDifference:%d\n", baseExp, power, baseExp-power)

    if (baseExp-power)<=30 {
        fmt.Printf("%s was caught\n", resp.Name)
        pokeDex.pokedexData[resp.Name] = pokeInfo{pokemonData:resp, caughtAt:time.Now().UTC()}
    } else {
        fmt.Println("catch failed")
    }
}

func commandInspect(pokemon string) {
    pokeInf, ok := pokeDex.pokedexData[pokemon]
    pokeData := pokeInf.pokemonData

    if !ok {
        fmt.Printf("You have not not caught %s\n", pokemon)
        return
    }

    fmt.Printf("Name: %s\n", pokeData.Name)
    //fmt.Printf("\nCaught at: %v\n", pokeInf.caughtAt)
    fmt.Printf("Height: %d\n", pokeData.Height)
    fmt.Printf("Weight: %d\n", pokeData.Weight)
    fmt.Println("Stats:")
    for _, stats := range pokeData.Stats {
        fmt.Printf("  -%s:%d\n", stats.Stat.Name, stats.BaseStat)
    }

    fmt.Println("Types:")
    for _, t := range pokeData.Types {
        fmt.Printf("  -%s\n", t.Type.Name)
    }
}

func commandPokedex(input string){
    fmt.Println("Your Pokedex:")
    for key, _ := range pokeDex.pokedexData {
        fmt.Printf(" -%s\n", key)
    }
}
