package main

import (
	"fmt"
	"math"
)

type Punto struct {
	x float64
	y float64
}

func main() {
	punto1 := Punto{0, 0}
	punto2 := Punto{10, 0}

	dist := distancia(punto1, punto2)
	fmt.Println(dist)
}

func distancia(p1 Punto, p2 Punto) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	result := math.Sqrt(x*x - y*y)
	return result
}
