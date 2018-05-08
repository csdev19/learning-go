package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		err1 := errors.New("error message")
		err2 := errors.New(12)
		err3 := errors.New(12.45)
		err4 := errors.New(true)
		fmt.Printf("%T\n", err1)
		fmt.Printf("%T\n", err2)
		fmt.Printf("%T\n", err3)
		fmt.Printf("%T\n", err4)
	*/
	resultado1, err1 := division(15, 2)
	resultado2, err2 := division(15, 0)

	fmt.Println(resultado1, err1)
	fmt.Println(resultado2, err2)
}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("no se puede dividir entre 0")
	} else {
		resultado = dividendo / divisor
	}
	return
}
