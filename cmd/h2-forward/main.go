package main

import (
	"log"

	"github.com/igarciaolaizola/h2-forward/internal/cli"
)

func main() {
	cmd := cli.NewCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
