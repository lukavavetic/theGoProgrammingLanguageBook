package main

import "fmt"

func main() {
	//x := 1
	//p := &x         // p, of type *int, points to x
	//fmt.Println(*p) // "1"
	//*p = 2          // equivalent to x = 2
	//fmt.Println(x)  // "2"
	//
	//name := "Luka"
	//
	//newName := &name
	//
	//fmt.Println(name) //Luka
	//fmt.Println(newName) //print address of variable name
	//fmt.Println(*newName) //print value of variable name - Luka
	//
	//*newName = "Matija"
	//
	//fmt.Println(name) // Matija
	//fmt.Println(*newName) // Matija
	//
	//var z, y int
	//fmt.Println(&z == &z, &z == &y, &z == nil) // "true false false"
	//
	//var h = funcName()
	//println(h) //return address of b variable

	v := 1
	incr(&v)              // side effect: v is now 2
	println(v)
	newResult := incr(&v)
	fmt.Println(newResult) // "3" (and v is 3)
	println(v)
}

//func funcName() *int {
//	b := 1
//
//	return &b
//}


func incr(p *int) int {
	println("pointer of param from func:", p) // adress of p parameter
	println("value of param from func:", *p) // value of p parameter

	*p++ // increments what p points to; does not change p
	return *p
}
