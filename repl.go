package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
)

func formatInput(input string) []string {
    input = strings.ToLower(input)
    input = strings.TrimSpace(input)
    inputArr := strings.Fields(input)
    return inputArr
}

func startRepl(){

    init_func()
    scanner := bufio.NewScanner(os.Stdin)

    commandList = map[string]cliCommand {
        "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
        },
        "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
        },
        "map": {
        name: "map",
        description: "Lists Location Areas",
        callback: commandMap,
        },
        "mapb": {
        name: "map",
        description: "Lists Location Areas",
        callback: commandMapBack,
        },
        "explore": {
        name: "explore <area-name>",
        description: "Lists Pokemons in the Area",
        callback: commandExplore,
        },
        "catch": {
        name: "catch <pokemon>",
        description: "Attempt to catch a Pokemon",
        callback: commandCatch,
        },
        "inspect": {
        name: "inspect <pokemon>",
        description: "Inspect a caught Pokemon",
        callback: commandInspect,
        },
        "pokedex": {
        name: "pokedex",
        description: "Lists all the Pokemons caught",
        callback: commandPokedex,
        },
    }
    for true {
        fmt.Printf("pokedex >")
        scanner.Scan()
        input := scanner.Text()
        inputArr := formatInput(input)
        if len(input)==0 {
            continue
        }
        cmd, ok := commandList[inputArr[0]]
        if ok {
            cmd.callback(inputArr[len(inputArr)-1])
        } else {
            fmt.Println("Invalid command")
        }
    }

}


