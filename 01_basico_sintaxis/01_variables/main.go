package main

import "fmt"

// Asi se hace un comentario

func main() {
	// crear esta variable
	var nombre string
	nombre = "cristian"
	// es lo mismo que hacer esto
	apellido := "sotomayor"
	// go deduce que tipo de variable es
	fmt.Println(nombre, apellido)
}
