package main

import "fmt"

func sumar(x, y int) int {
	return x + y
}

func esMenor(edad uint8) bool {
	if edad < 20 {
		return true
	} else {
		return false
	}
}

func main() {
	// funciones
	//suma := sumar(3, 2)
	//fmt.Println(suma)

	//numero, ok := isPositive(123)
	//fmt.Println(numero, ok)

	//manyValues(1, 3, 124, 125, 126, 6, 326, 23, 632, 732, 74)
	valor := someValue()
	fmt.Println(valor)
}

func isPositive(integer int) (int, bool) {
	if integer > 0 {
		return integer, true
	} else {
		return integer, false
	}
}

func manyValues(values ...int) {
	fmt.Print(values)
	fmt.Printf("%T\n", values)
}

func someValue() (result int) {
	result = 123
	return
}
