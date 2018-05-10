package main

import (
	"fmt"
	"math"
)

/*
	Aqui pondremos un ejemplo de interfaces con figuras
	geometricas. AREA Y PERIMETROS
*/

// Estructura e interface CIRCULO
type Circulo struct {
	radio float64
}

func (c Circulo) area() (resultado float64) {
	resultado = math.Pi * (c.radio * c.radio)
	return resultado
}

func (c Circulo) perimetro() (resultado float64) {
	resultado = math.Pi * c.radio * 2
	return resultado
}

// Estructura e interface CUADRADO
type Cuadrado struct {
	lado float64
}

func (c Cuadrado) area() (resultado float64) {
	resultado = c.lado * c.lado
	return resultado
}

func (c Cuadrado) perimetro() (resultado float64) {
	resultado = c.lado * 4
	return resultado
}

// Estructura e interface RECTANGULO
type Rectangulo struct {
	base   float64
	altura float64
}

func (r Rectangulo) area() (resultado float64) {
	resultado = r.base * r.altura
	return resultado
}

func (r Rectangulo) perimetro() (resultado float64) {
	resultado = r.base*2 + r.altura
	return resultado
}

// Interface Figuras
type Figuras interface {
	area() float64
	perimetro() float64
}

func main() {
	// Interfaces
	circulo := Circulo{10.0}
	cuadrado := Cuadrado{10}
	rectangulo := Rectangulo{5.0, 2.0}

	area2 := area(rectangulo)
	area := area(cuadrado)
	perimetro := perimetro(circulo)
	fmt.Println("el area es: ", area)          // 100
	fmt.Println("el area es: ", area2)         // 100
	fmt.Println("el perimetro es:", perimetro) // 62,..
}

func area(f Figuras) float64 {
	return f.area()
}

func perimetro(f Figuras) float64 {
	return f.perimetro()
}
