package main

import (
	"math"
	"fmt"
)

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}
}
