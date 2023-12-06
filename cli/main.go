package main

import (
	"os"

	"github.com/avarian/online-shopping-cart/cli/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
