package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

func (p *Persona) darEdad(edad int8) {
	p.edad = edad
}

func (p Persona) saludar() {
	fmt.Println("hola soy una persona")
}

func main() {
	var persona Persona

	persona.saludar()

	fmt.Println(persona.edad)
	persona.darEdad(20)
	fmt.Println(persona.edad)
}
