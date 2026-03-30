package main

import (
	"log"

	"example.com/cuit-go-frameworks/cobra-demo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
