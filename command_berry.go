package main

import (
	"fmt"
	"errors"
)

func commandBerry(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("error: berry name required")
	}

	berryName := args[0]
	berry, err := cfg.pokeapiClient.GetBerry(berryName)
	if err != nil {
		return err
	}

	fmt.Println("Name: ", berry.Name)
	fmt.Println("Growth Time: ", berry.GrowthTime)
	fmt.Println("Max Harvet: ", berry.MaxHarvest)
	fmt.Println("Natural Gift Power: ", berry.NaturalGiftPower)
	fmt.Println("Natural Gift Type: ", berry.NaturalGiftType.Name)
	fmt.Println("Size: ", berry.Size)
	fmt.Println("Smoothness: ", berry.Smoothness)
	fmt.Println("SoilDryness: ", berry.SoilDryness)
	fmt.Println("Firmness: ", berry.Firmness.Name)
	fmt.Println("Item Name: ", berry.Item.Name)

	fmt.Println("Flavors: ")
	for _, flavor := range berry.Flavors {
		fmt.Println("	-", flavor.Flavor.Name)
		fmt.Println("	  -Potency: ", flavor.Potency)
	}

	

	return nil
}