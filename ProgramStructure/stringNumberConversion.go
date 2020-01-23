package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Println(y, strconv.Itoa(x)) //123 123

	fmt.Println(strconv.FormatInt(int64(x), 2)) //1111011

	s := fmt.Sprintf("x=%b", x)

	println(s) //x=1111011

	x, err := strconv.Atoi("123")

	if err != nil {

	}

	z, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

	if err != nil {

	}

	println(z)
}
