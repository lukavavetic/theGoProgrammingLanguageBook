package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//r := basename("file.go")
	//	//println("\n", r)

	comma("123456789")


	//s := "abc"
	//b := []byte(s)
	//println(b) //[3/32]0xc00007ef20
	//s2 := string(b)
	//println(s2) //abc


	println(intsToString([]int{1, 2, 3}))

}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
//func basename(s string) string {
//	// Discard last '/' and everything before.
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '/' {
//			s = s[i+1:]
//			break
//		}
//	}
//
//	// Preserve everything before last '.'.
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '.' {
//			s = s[:i]
//			break
//		}
//	}
//	return s
//}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	println(slash)
	s = s[slash+1:]
	print(s)
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	//println(s)
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')

	return buf.String()
}