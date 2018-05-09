package main

import (
	"fmt"

	"github.com/cris/paquetes/example"
)

func main() {
	// Vamos a usar lo que esta dentro de example
	example.Example()
	example.Num_test = 123
	fmt.Println(example.Num_test)
}
