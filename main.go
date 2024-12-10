package main

import (
	"os"

	"github.com/OctaneAL/Shortly/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
