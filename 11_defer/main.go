package main

import "fmt"

func main() {
	// Defer
	i := 0
	for i < 1000 {
		i++
		if i == 999 {
			fmt.Println("primer conteo hasta", i)
		}
	}
	defer endCount(i)

	i = 0
	for i < 5000 {
		i++
		if i == 4999 {
			fmt.Println("contando hasta", i)
		}
	}
	defer endCount2(i)
}

func endCount(contador int) {
	contador--
	fmt.Println("Se finalizo el primer conteo hasta", contador)
}

func endCount2(contador int) {
	contador--
	fmt.Println("se finalizo el segundo conteo hasta", contador)
}

/*
el resultado
	primer conteo hasta 999
	contando hasta 4999
	se finalizo el segundo conteo hasta 4999
	Se finalizo el primer conteo hasta 999
*/
