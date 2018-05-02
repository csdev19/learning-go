package main

import "fmt"

func main() {
	// Maps -> diccionarios
	dias := make(map[string]string) // claves string y valor string
	dias["lunes"] = "Lunes"
	dias["martes"] = "Martes"
	dias["miercoles"] = "Miercoles"
	fmt.Println(dias)
	// Otra forma
	dias2 := map[string]string{"lunes": "Lunes", "martes": "Martes", "miercoles": "Miercoles"}
	fmt.Println(dias2)

	delete(dias2, "martes")
	fmt.Println(dias2)

	if dia1, ok := dias["jueves"]; ok {
		fmt.Println(dia1, "existe y es", ok)
	} else {
		fmt.Println(dia1, "no existe y es", ok)
		fmt.Printf("%T\n", dia1)
	}
}
