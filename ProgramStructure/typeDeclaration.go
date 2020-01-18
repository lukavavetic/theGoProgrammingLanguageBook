package main

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func main() {
	result := CToF(23)

	println(result)
}

func CToF(c Celsius) Fahrenheit
{
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius
{
	return Celsius((f - 32) * 5 / 9)
}


//var c Celsius
//var f Fahrenheit
//fmt.Println(c == 0) // "true"
//fmt.Println(f >= 0) // "true"
//fmt.Println(c == f) // compile error: type mismatch
//fmt.Println(c == Celsius(f)) // "true"!


