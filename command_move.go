package main

import (
	"fmt"
	"errors"
)

func commandMove(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("error: you must provide a move name")
	}

	name := args[0]
	move, err := cfg.pokeapiClient.GetMove(name)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s:\n", move.Name)
	fmt.Println("Accuracy: ", move.Accuracy)
	fmt.Println("Power: ", move.Power)
	fmt.Println("PP: ", move.Pp)
	fmt.Println("Effect Chance: ", move.EffectChance)
	return nil
}