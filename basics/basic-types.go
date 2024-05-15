package main

import "fmt"

func main() {
	a := 2
	b := 3.1
	fmt.Printf("a is of Type: %T and Value: %[1]v\n", a)
	fmt.Printf("b is of Type: %T and Value: %[1]v\n", b)

	a = int(b)
	fmt.Printf("a is of Type: %T and Value: %[1]v\n", a)

	// a = b // This will not work as Go is a statically typed language
	b = float64(a)
	fmt.Printf("b is of Type: %T and Value: %[1]v\n", b)
}
