package main

import (
	"fmt"
	//"unicode/utf8"
)


func main() {
	//s := "Hello, BF"
	//fmt.Println(len(s)) // "9"
	//fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i, r := range "Hello, BF" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println(string(65)) // "A", not "65"
	fmt.Println(string(1234567)) // "?"
}