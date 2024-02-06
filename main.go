package main

import (
	"github.com/ismailash/be-enigma-laundry/delivery"
)

func main() {
	// Run app
	delivery.NewServer().Run()
}
