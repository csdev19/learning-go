package main

/*
  importamos reflect para llamar a su metodo
  TypeOf() para inspeccionar el tipo de las variables
*/

import (
	"fmt"
	"reflect"
)

func main() {
	// y aqui podemos probar pasandole valores
	// a variables y que go detecte que tipo de datos son
	nombre := "cristian"
	numeroEntero := 123
	numeroDecimal := 10.41
	booleano := false

	fmt.Println(reflect.TypeOf(nombre))
	fmt.Println(reflect.TypeOf(numeroEntero))
	fmt.Println(reflect.TypeOf(numeroDecimal))
	fmt.Println(reflect.TypeOf(booleano))
	// forma alternativa
	fmt.Printf("nombres es de tipo: %T\n", nombre)
	fmt.Printf("nombres es de tipo: %T\n", numeroEntero)
	fmt.Printf("nombres es de tipo: %T\n", numeroDecimal)
	fmt.Printf("nombres es de tipo: %T\n", booleano)
	// valor cero
	var nombre1 string
	// nombre = ""
	fmt.Println(nombre1) //-> retorna linea en blanco
	var numero int
	fmt.Println(numero) //-> retorna linea en blanco
	// numero = 0
	var entiendes bool
	entiendes = false //OJO!!!!!
	fmt.Println(entiendes)

}
