package main

import "fmt"

func main() {
	imprimir := func() {
		fmt.Println("hola")
	}
	fmt.Printf("%T\n", imprimir)

}
