package main

import (
	"os"

	"github.com/avarian/online-shopping-cart/cli/commands"
	_ "github.com/avarian/online-shopping-cart/docs"
)

// @title           Online Shopping Cart API
// @version         1.0
// @description     This is a API for Online Shopping Cart API.

// @host localhost:8080
func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
