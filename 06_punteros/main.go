package main

import "fmt"

func main() {
	// Punteros
	a := 12
	// devuelve la direccion en memoria
	fmt.Println(&a)

	// &a devuelve *int
	b := &a

	fmt.Printf("%T\n", b)

	x := 10

	fmt.Println("x es :", x)
	changeToZero(&x)
	fmt.Println("x es :", x)
}

func changeToZero(numero *int) {
	*numero = 0
}
