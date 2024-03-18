package main

import (
	"fmt"
)

func commandBerries(cfg *config, args ...string) error {
	berries, err := cfg.pokeapiClient.ListBerries(cfg.nextBerriesURL)
	if err != nil {
		return err
	}

	cfg.nextBerriesURL = berries.Next
	cfg.prevBerriesURL = berries.Previous

	for _, berry := range berries.Results {
		fmt.Println(berry.Name)
	}
	return nil
}

func commandBerriesb(cfg *config, args ...string) error {
	if cfg.prevBerriesURL == nil {
		return errors.New("you're on the first page")
	}

	berries, err := cfg.pokeapiClient.ListBerries(cfg.prevBerriesURL)
	if err != nil {
		return err
	}

	cfg.nextBerriesURL = berries.Next
	cfg.prevBerriesURL = berries.Previous

	for _, berry := range berries.Results {
		fmt.Println(berry.Name)
	}
	return nil
}