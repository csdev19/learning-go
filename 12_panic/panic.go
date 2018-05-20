package main

import "fmt"

func main() {
	// Panic
	division(45, 0)

}

func division(dividendo, divisor int) {
	// para que imprima luego de la ejecucion
	defer fmt.Println("esto se ejecutara al final y aunque salte un error, pase lo que pase")
	if divisor == 0 {
		// le podemos pasar informacion para que se imprima en la terminal
		panic("ocurrio una division entre cero")
	}
	fmt.Println(dividendo / divisor)
}
