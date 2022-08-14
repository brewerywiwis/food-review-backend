package main

import (
	"fmt"

	"github.com/brewerywiwis/food-review-backend/pkg/interface/rest"
)

func main() {
	fmt.Println("HELLO WORLD")
	server := rest.NewServer()

	server.Run(":8090")
}
