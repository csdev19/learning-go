package main

import "fmt"

type Point struct {
	point_x int
	point_y int
	point_z int
}

func main() {
	// Estructuras:
	mini := Point{2, 3, 4}
	//fmt.Println(mini)
	//fmt.Printf("%T\n", mini)
	//mini.point_x = 5
	//fmt.Println(mini)
	mini2 := Point{
		point_x: 2,
		point_y: 3,
		point_z: 4,
	}
	mini3 := Point{
		point_y: 3,
		point_z: 4,
		point_x: 2,
	}
	fmt.Println(mini)
	fmt.Println(mini2)
	fmt.Println(mini3)
}
