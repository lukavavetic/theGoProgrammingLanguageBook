package main

func main() {
	//x = 1 //named variable
	//*p = true // indirect variable
	//person.name = "bob" // struct field
	//count[x] = count[x] * scale // array or slice or map element count[x] *= scale

	//i, j, k = 2, 3, 5

	fib(5)
}


func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		println("value x", x) //0 for first iteration
		println("value y", y) // 1 for first iteration
	x, y = y, x+y
		println("value x", x) // 1 for first iteration
		println("value y", y) // 1 for first iteration
	break
	}

	return x
}