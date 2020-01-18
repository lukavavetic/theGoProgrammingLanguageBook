package main

import "fmt"

func main() {
	//var x uint8 = 1<<1 | 1<<5
	//	//var y uint8 = 1<<1 | 1<<2
	//	//fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	//	//fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	//	//fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	//	//fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	//	//fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	//	//fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	//	//
	//	//for i := uint(0); i < 8; i++ {
	//	//	if x&(1<<i) != 0 { // membership test
	//	//		fmt.Println(i) // "1", "5"
	//	//	}
	//	//}
	//	//
	//	//fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	//	//fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
	//	//
	//	//var apples int32 = 1
	//	//var oranges int16 = 2
	//	////var compote int = apples + oranges // compile error invalid operation: apples + oranges (mismatched types int32 and int16)
	//	//var compote = int(apples) + int(oranges)
	//	//fmt.Println(compote)
	//	//
	//	//
	//	//
	//	//f := 3.141 // a float64
	//	//i := int(f)
	//	//fmt.Println(f, i)   // "3.141 3"
	//	//f = 1.99
	//	//fmt.Println(int(f)) // "1"

	//f := 1e100 // a float64
	//	//i := int(f) // result is implementation-dependent
	//	//fmt.Printf("%d\t%f\n", i, f)

	//o := 0666
	//fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	//x := int64(0xdeadbeef)
	//fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'A'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
}

