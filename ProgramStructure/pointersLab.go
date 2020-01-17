package main
import (
	"flag"
	"fmt"
	"strings"
)
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	fmt.Println(n) //address of n
	fmt.Println(sep) //address of n
	fmt.Println(*n) //value of n = false
	fmt.Println(*sep) //value of sep

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	} }