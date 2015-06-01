package main

import (
	"ccode"
	"fmt"
)

func main() {
	if !ccode.True() {
		fmt.Printf("True(): expected %v got %v\r\n", true, ccode.True())
	}

	a, b := 42, -6
	if ccode.Max(a, b) != a {
		fmt.Printf("Max(%v, %v): expected %v, got %v\r\n", a, b, a, ccode.Max(a, b))
	}

	v := 9000
	if ccode.Inc(&v); v < 9001 {
		fmt.Printf("Inc(9000): expected 9001, got %v\r\n", v)
	}
}
