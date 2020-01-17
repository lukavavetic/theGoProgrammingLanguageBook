package main

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	name := "Luka"

	newName := &name

	fmt.Println(name) //Luka
	fmt.Println(newName) //print address of variable name
	fmt.Println(*newName) //print value of variable name - Luka

	*newName = "Matija"

	fmt.Println(name) // Matija
	fmt.Println(*newName) // Matija

	var z, y int
	fmt.Println(&z == &z, &z == &y, &z == nil) // "true false false"
}
