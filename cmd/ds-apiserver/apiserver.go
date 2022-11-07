package main

import (
	"io.github/devopssphere/cmd/ds-apiserver/app"
	"log"
)

func main() {
	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
