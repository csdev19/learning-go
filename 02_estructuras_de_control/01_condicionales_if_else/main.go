package main

import "fmt"

/*
Simbolos:
-> || Or
-> && And
-> !  Not
-> <  menor que
-> >  mayor que
-> <= menor igual que
-> >= mayor igual que
-> == igual que
-> != diferente de
*/

func main() {
	//isValid1 := true
	if isValid := true; isValid {
		fmt.Printf("%T\n", isValid)
		fmt.Println("isValid es verdadero")
	} else {
		fmt.Println("no es verdadero")
	}

	age := 12

	if nombre := "Cristian"; age < 8 {
		fmt.Println(nombre, "tiene menos de 8")
	} else if age < 16 {
		fmt.Println(nombre, "tiene menos de 16")
	} else if age < 25 {
		fmt.Println(nombre, "tiene menos de 25")
	} else {
		fmt.Println(nombre, "tiene mas de 25")
	}

	fmt.Println("fin del programa")
}
