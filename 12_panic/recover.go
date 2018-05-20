package main

import "fmt"

func main() {
	division(29, 0)
}

func division(dividendo, divisor int) {
	defer func() {
		// r es lo que devuelve panic
		if r := recover(); r != nil {
			fmt.Println("recover recibio un tipo string")
		}
	}()

	if divisor == 0 {
		panic("dividio entre cero")
	}

	fmt.Println("se ejecuto normal")
}
