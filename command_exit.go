package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Thank you for using PokedexCLI!")
	fmt.Println("We hope to see you again soon.")
	fmt.Println("Goodbye 👋")
	fmt.Println()
	os.Exit(0)
	return nil
}