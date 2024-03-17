package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("error: you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catch_success := true

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catch_success {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	}else{
		fmt.Printf("%s escaped!", pokemon.Name)
	}
	return nil
}