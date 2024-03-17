package commands

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("Your PokeDex is empty. Catch Pokemon to learn about them.")
		return nil
	}

	fmt.Println("Your PokeDex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("	-", pokemon.Name)
	}
	return nil
}