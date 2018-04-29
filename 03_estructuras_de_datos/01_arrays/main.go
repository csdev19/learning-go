package main

import "fmt"

func main() {
	// Arrays
	var array1 [10]string
	array1[0] = "inicio"
	array1[1] = "Cristian"
	array1[2] = "Diego"
	array1[3] = "Maribel"
	fmt.Printf("%T\n", array1)
	fmt.Println(array1)
	// Otra forma de crear arrays
	array2 := [10]string{"inicio", "cristian", "luis"}
	fmt.Printf("%T\n", array2)
	fmt.Println(array2)
	// ver el tamaño del array
	// LEN
	size := len(array2)
	fmt.Println("el tamaño es", size)
	// obtener partes del array
	// array[inicio:(final-1)]
	parte1 := array2[0:2]
	fmt.Println(parte1)
	fmt.Println(array2[0:2])
	// valores vacios
	var empty_int [3]int
	var empty_string [3]string
	var empty_float [3]float32

	fmt.Println(empty_int)
	fmt.Println(empty_string)
	fmt.Println(empty_float)
}
