package main

import "fmt"

func main() {
	// Slices
	var slices1 []string
	slices1 = append(slices1, "Cristian")
	fmt.Printf("el tamano es: %d y su capacidad: %d\n", len(slices1), cap(slices1))
	slices1 = append(slices1, "Diego")
	fmt.Printf("el tamano es: %d y su capacidad: %d\n", len(slices1), cap(slices1))
	slices1 = append(slices1, "Maribel")
	fmt.Printf("el tamano es: %d y su capacidad: %d\n", len(slices1), cap(slices1))

	// luego de ver como las capacidades de memoria de
	// los slices se duplicaban pudimos comprobar que
	// reglas siguen

	fmt.Printf("%T\n", slices1)
	slices1 = append(slices1, "Luis")
	fmt.Println(slices1)
	slices1[2] = "Luis"
	slices1 = append(slices1, "Maribel")
	fmt.Println(slices1)
	// Forma comun de crear slices
	slices2 := make([]string, 0)
	fmt.Printf("%T\n", slices2)

	// Agregar valores de manera inmediata
	slices3 := []string{
		"Cristian",
		"Diego",
		"Maribel",
		"Luis",
	}
	fmt.Printf("%T\n", slices3)
	fmt.Println("fin del programa")
}
