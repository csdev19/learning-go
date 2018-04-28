package main

import "fmt"

func main() {
	a := 3
	dia := 4
	switch a {
	case 1:
		fmt.Println("el valor es 1")
	case 2:
		fmt.Println("el valor es 2")
	case 3:
		fmt.Println("el valor es 3")
	default:
		fmt.Println("el valor es mayor a 3")
	}

	// si queremos un comportamiento que daria el break

	// o podemos hacer lo siguiente
	/*switch dia {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas en un dia de la semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Es fin de la semana")
	}
	default:
		fmt.Println("no es un dia valido")
	*/
	switch dia2 := 2; {
	case a >= 0 && a <= 5:
		fmt.Println("Estas entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("no es un dia de semana")
	}
}
