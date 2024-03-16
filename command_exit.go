package main

import (
	"os"
	"fmt"
)

func commandExit(cfg *config) error {
	fmt.Println()
	fmt.Println("Thank you for using PokedexCLI!")
	fmt.Println("We hope to see you again soon.")
	fmt.Println("Goodbye 👋")
	fmt.Println()
	os.Exit(0)
	return nil
}