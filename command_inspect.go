package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("error: you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("error: you have not caught this pokemon yet")
	}

	
	
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Moves:")
	for _, move := range pokemon.Moves {
		fmt.Printf("	-%s\n", move.Move.Name)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("	-", typeInfo.Type.Name)
	}
	return nil
}